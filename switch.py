import glob
import os
import re
import ssl
import subprocess
import sys
import json
import zipfile
import shutil
import shlex
from collections import defaultdict

try:
    from urllib.request import Request, urlopen, urlretrieve
    from io import BytesIO # Python 3
except ImportError:
    from urllib2 import Request, urlopen, urlretrieve
    from StringIO import StringIO as BytesIO # Python 2

ssl._create_default_https_context = ssl._create_unverified_context

INSTALL_DIR = os.path.expanduser("~/.odin")
OLD_ODIN = "/usr/local/bin/odin"


def find_odin_file(directory=INSTALL_DIR, prefix="odin-"):
    pattern = os.path.join(directory, f"{prefix}*")  # Match files like 'odin-*'
    files = glob.glob(pattern)
    if files:
        return files[0]
    return OLD_ODIN


NEW_ODIN = find_odin_file()

branch = "update/mig-script"
odin_backend_address = "http://odin-backend.d11dev.com"
odin_access_key = ''
odin_secret_access_key = ''
odin_access_token = ''
envCheckUri = "/api/integration/cli/v2/deployerint/envcheck"
checkMigrationStatusUri = "/api/integration/cli/v2/deployerint/migrationstatus"


def is_dreampay():
    config_path = os.path.expanduser("~/.odin/config")
    try:
        with open(config_path, "r") as file:
            for line in file:
                if line.strip().startswith("backend_addr:"):
                    backend_address = line.split(":", 1)[1].strip()
                    return "dreampay" in backend_address
    except FileNotFoundError as e:
        print("Config file not found, reading from env variables")
        return None

    return os.getenv("ODIN_BACKEND_ADDRESS") and "dreampay" in os.getenv("ODIN_BACKEND_ADDRESS")


def set_tokens(config_file):
    global odin_access_key, odin_secret_access_key, odin_access_token
    if os.path.isfile(config_file):
        with open(config_file, "r") as file:
            for line in file:
                line = line.strip()  # Remove leading/trailing whitespace
                if line.startswith("access_key"):
                    if ':' in line:
                        odin_access_key = line.split(":", 1)[1].strip()
                    if '=' in line:
                        odin_access_key = line.split("=", 1)[1].strip()
                elif line.startswith("secret_access_key"):
                    if ':' in line:
                        odin_secret_access_key = line.split(":", 1)[1].strip()
                    if '=' in line:
                        odin_secret_access_key = line.split("=", 1)[1].strip()
                elif line.startswith("access_token"):
                    if ':' in line:
                        odin_access_token = line.split(":", 1)[1].strip()
                    if '=' in line:
                        odin_access_token = line.split("=", 1)[1].strip()
    if not odin_access_key or not odin_secret_access_key:
        print("Config file not found. Enter keys.")
        odin_access_key = input("Enter ODIN_ACCESS_KEY: ").strip()
        odin_secret_access_key = input("Enter ODIN_SECRET_ACCESS_KEY: ").strip()
        # Call configure for both old and new odin
        configure()
        exit(0)


def get_current_bin_version():
    try:
        files = os.listdir(INSTALL_DIR)
        for file in files:
            if file.startswith('odin-'):
                return file.split("-")[1]
    except FileNotFoundError:
        return None

def get_env_from_config(config_path):
    try:
        with open(config_path, "r") as file:
            for line in file:
                if line.strip().startswith("envName:"):
                    return line.split(":", 1)[1].strip()
    except FileNotFoundError as e:
        print(f"Error reading config file {config_path}: {e}")
        return None

def process_env_argument():
    excluded_verbs = {"configure", "list", "create", "set", "version", "update"}
    if any(verb in sys.argv for verb in excluded_verbs):
        return

    if "--env" or "--name" not in sys.argv:
        # if no --env or --name provided, read from config
        config_file = os.path.expanduser("~/.odin/config")
        env_name = get_env_from_config(config_file)
        if env_name:
            return env_name
        else:
            return None
    else:
        if "--env" in sys.argv:
            return sys.argv[sys.argv.index("--env") + 1]
        elif "--name" in sys.argv:
            return sys.argv[sys.argv.index("--name") + 1]


def update_binary():
    version_url = "https://artifactory.dream11.com/migrarts/odin-artifact/odin-version.txt"

    try:
        response = urlopen(version_url)

        if response.getcode() == 200:
            latest_version = response.read().decode("utf-8").strip()  # Decode for compatibility

            current_version = get_current_bin_version()

            if not latest_version and current_version is not None:
                print("Error: The version fetched from {} is empty.".format(version_url))
                return

            if current_version is None or current_version < latest_version:
                print("Updating odin binary to version {}".format(latest_version))

                # Step 3: Download the binary zip from Artifactory
                binary_url = "https://artifactory.dream11.com/migrarts/odin-artifact/odin-artifact.zip"
                zip_response = urlopen(binary_url)

                if zip_response.getcode() == 200:
                    zip_content = zip_response.read()
                    with zipfile.ZipFile(BytesIO(zip_content)) as zip_ref:
                        zip_ref.extractall(INSTALL_DIR)

                    extracted_folder = os.path.join(INSTALL_DIR, "cli-migration")
                    binary_filepath = os.path.join(extracted_folder, "odin-{}".format(latest_version))
                    final_binary_path = os.path.join(INSTALL_DIR, "odin-{}".format(latest_version))

                    if os.path.exists(binary_filepath):
                        os.rename(binary_filepath, final_binary_path)
                        os.chmod(final_binary_path, 0o755)

                        shutil.rmtree(extracted_folder, ignore_errors=True)

                        subprocess.call("sudo spctl --master-enable", shell=True)
                        subprocess.call('xattr -dr com.apple.quarantine "{}"'.format(final_binary_path), shell=True)

                        print("Successfully updated to version {}.".format(latest_version))
                    else:
                        print("Error: The binary {} was not found after extraction.".format(binary_filepath))
                else:
                    print("Error: Failed to download the binary zip. Status code: {}".format(zip_response.getcode()))
            else:
                print("Binary up-to-date with version {}".format(current_version))
        else:
            print("Error: Failed to fetch version file from Artifactory. Status code: {}".format(response.getcode()))

    except Exception as e:
        print("Error: Unexpected error occurred: {}".format(e))


def execute_new_odin():
    subprocess.call([NEW_ODIN] + sys.argv[1:])
    exit(0)


def execute_new_odin_with_custom_cmd(arg_list):
    subprocess.call([NEW_ODIN] + arg_list)

def execute_old_odin():
    subprocess.call([OLD_ODIN] + sys.argv[1:])
    exit(0)


def configure():
    global odin_access_key, odin_secret_access_key, odin_backend_address
    # Set environment variable for new odin and configure
    new_odin_env = os.environ.copy()

    if os.getenv("ODIN_BACKEND_ADDRESS"):
        odin_backend_address = os.getenv("ODIN_BACKEND_ADDRESS")
    new_odin_env["ODIN_BACKEND_ADDRESS"] = odin_backend_address.strip('"').strip("'")
    new_odin_env["ODIN_ACCESS_KEY"] = odin_access_key.strip('"').strip("'")
    new_odin_env["ODIN_SECRET_ACCESS_KEY"] = odin_secret_access_key.strip('"').strip("'")

    print("Configuring old odin")
    subprocess.call([OLD_ODIN, "configure"], env=new_odin_env)

    if not is_dreampay():
        print("Configuring new odin")
        new_odin_env["ODIN_BACKEND_ADDRESS"] = "odin-deployer.dss-platform.private:80"
        subprocess.call([NEW_ODIN, "configure"], env=new_odin_env)


def check_env_exists_in_old_odin(env_name):
    global odin_access_token, odin_backend_address, envCheckUri
    url = odin_backend_address + envCheckUri + "/?env_name=" + env_name
    req = Request(url)
    req.add_header('Authorization', 'Bearer ' + odin_access_token)
    req.add_header('App-Version', '1.4.1')
    req.add_header('Accept', 'application/json')
    try:
        content = urlopen(req).read()
        content_json = json.loads(content)
        return content_json.get("status") == 200
    except Exception:
        return False


def is_service_migrated_to_new_odin(service_name, env_name):
    global odin_access_token, odin_backend_address, checkMigrationStatusUri
    url = odin_backend_address + checkMigrationStatusUri + "/" + env_name + "/" + service_name
    req = Request(url)
    req.add_header('Authorization', 'Bearer ' + odin_access_token)
    req.add_header('App-Version', '1.4.1')
    req.add_header('Accept', 'application/json')
    try:
        content = urlopen(req).read()
        content_json = json.loads(content)
        return content_json.get("migrationStatus") == "SUCCESS"
    except Exception:
        return False

def get_service_name_from_file(file_path):
    try:
        with open(file_path, 'r') as f:
            data = json.load(f)
            return data.get("name")
    except (FileNotFoundError, json.JSONDecodeError) as e:
        print(f"Error reading file {file_path}: {e}")
        sys.exit(1)

def display_all_envs(old_env_list, new_env_list):
    envs = []
    # Decode and remove ANSI escape sequences
    clean_old_envs = re.sub(r'\x1b\[[0-9;]*m', '', old_env_list.decode())

    # Extract rows containing environment details
    rows = clean_old_envs.split("\n")[3:]  # Skip headers and separator lines

    # Parse NAME, TEAM, CREATED BY, ENV TYPE, STATE, ACCOUNT
    for row in rows:
        env_data = row.split("|")
        if len(env_data) > 4:
            name, env_type, state, account = env_data[0], env_data[3], env_data[4], env_data[5]
            envs.append({
                "name": name.strip(),
                "env_type": env_type.strip(),
                "state": state.strip(),
                "account": account.strip(),
            })

    # Decode and remove ANSI escape sequences
    clean_new_envs = re.sub(r'\x1b\[[0-9;]*m', '', new_env_list.decode())

    # Extract rows containing environment details
    rows = clean_new_envs.split("\n")[2:]  # Skip header lines

    # Parse name, state, and account
    for row in rows:
        env_data = row.split("|")
        if len(env_data) > 2:
            name, state, account = env_data[0], env_data[1], env_data[2]
            envs.append({"name": name.strip(), "state": state.strip(), "account": account.strip()})

    # merge envs with same name but different accounts
    # Group by name and combine accounts
    grouped_envs = defaultdict(lambda: {"env_type": "", "state": "", "accounts": set()})

    for env in envs:
        key = env["name"].strip() or "N/A"
        grouped_envs[key]["env_type"] = env.get("env_type", "").strip() or "N/A"
        grouped_envs[key]["state"] = env.get("state", "").strip() or "N/A"
        if env["account"].strip():
            grouped_envs[key]["accounts"].add(env["account"].strip())  # Use set to remove duplicates

    # Convert accounts set to a comma-separated string (or "N/A" if empty)
    for key in grouped_envs:
        grouped_envs[key]["accounts"] = ", ".join(sorted(grouped_envs[key]["accounts"])) if grouped_envs[key]["accounts"] else "N/A"

    # Define column headers
    headers = ["NAME", "ENV TYPE", "STATE", "ACCOUNTS"]

    # Determine column widths dynamically
    col_widths = [
        max(len(key) for key in list(grouped_envs.keys()) + [headers[0]]),
        max(len(row["env_type"]) for row in list(grouped_envs.values()) + [{"env_type": headers[1]}]),
        max(len(row["state"]) for row in list(grouped_envs.values()) + [{"state": headers[2]}]),
        max(len(row["accounts"]) for row in list(grouped_envs.values()) + [{"accounts": headers[3]}]),
    ]

    # Print the header row
    header_format = " | ".join(f"{{:<{w}}}" for w in col_widths)
    print(header_format.format(*headers))
    print("-" * (sum(col_widths) + len(col_widths) * 3 - 3))  # Print separator

    # Print each grouped row
    for name, details in grouped_envs.items():
        print(header_format.format(name, details["env_type"], details["state"], details["accounts"]))


def transform_service_set_file(content):
    try:
        data = json.loads(content)
    except ValueError:
        print("Error: Not a valid JSON file")
        sys.exit(1)

    if not isinstance(data, dict) or 'services' not in data or not isinstance(data['services'], list):
        print("Error: JSON file must contain a 'services' array")
        sys.exit(1)

    # Transform each service in the array
    for service in data['services']:
        if isinstance(service, dict) and 'version' in service:
            version = service['version']

            if version == 'stable':
                del service['version']
                service['labels'] = 'isStable=true'

            elif version == 'dev-stable':
                del service['version']
                service['labels'] = 'isDevStable=true'

            elif version == 'load-stable':
                del service['version']
                service['labels'] = 'isLoadStable=true'

    return json.dumps(data, indent=2)

def create_new_service_set_and_trigger_odin(original_file):
    if os.path.exists(original_file):
        with open(original_file, "r") as f:
            content = f.read()

        transformed_content = transform_service_set_file(content)
        new_directory = os.path.expanduser("~/.odin/tmp/service-sets")
        os.makedirs(new_directory, exist_ok=True)
        base_file = os.path.basename(original_file)
        new_filename = os.path.join(new_directory, base_file.replace(".json", "-new-odin.json"))
        with open(new_filename, "w") as f:
            f.write(transformed_content)

        updated_args = sys.argv[1:]
        file_index = updated_args.index("--file") + 1
        updated_args[file_index] = new_filename
        execute_new_odin_with_custom_cmd(updated_args)

def main():
    global odin_access_key, odin_secret_access_key, odin_access_token, odin_backend_address, OLD_ODIN
    env_name = None
    # If /opt/homebrew/bin/odin exists, use this as old_odin
    if os.path.isfile("/opt/homebrew/bin/odin"):
        OLD_ODIN = "/opt/homebrew/bin/odin"
    env_name = process_env_argument()
    if len(sys.argv) == 1:
        execute_new_odin()

    elif len(sys.argv) == 2 and "--version" in sys.argv:
        execute_old_odin()

    elif len(sys.argv) == 2 and "version" in sys.argv:
        execute_new_odin()

    elif "configure" in sys.argv:
        config_file = os.path.expanduser("~/.odin/config")

        # Delete ~/.odin/config
        if os.path.isfile(config_file):
            os.remove(config_file)

        # Call configure for both old and new odin
        configure()

        exit(0)

    elif is_dreampay():
        execute_old_odin()

    elif "environment" in sys.argv and "operate" in sys.argv:
        print("Command not available")
        exit(0)

    elif "label" in sys.argv:
        print("Label commands have been deprecated")
        execute_new_odin()

    elif "service-set" in sys.argv:
        if "list" in sys.argv or "describe" in sys.argv:
            print("Command deprecated, refer to the service set file on https://github.com/dream11/service-sets to learn more about it")
            exit(0)
        elif "create" in sys.argv or "delete" in sys.argv:
            print("Command deprecated, use file to deploy")
            exit(0)
        elif "--env" in sys.argv:
            env_name = sys.argv[sys.argv.index("--env") + 1]
            if check_env_exists_in_old_odin(env_name):
                execute_old_odin()
            else:
                if "--file" in sys.argv:
                    create_new_service_set_and_trigger_odin(sys.argv[sys.argv.index("--file") + 1])
        else:
            execute_new_odin()

    elif "env" not in sys.argv and "--env" not in sys.argv:
        if "list" in sys.argv:
            if "service" in sys.argv:
                execute_new_odin()
            else:
                execute_old_odin()

        if "describe" in sys.argv:
            if "service" in sys.argv:
                execute_new_odin()
            else:
                execute_old_odin()

        if "release" in sys.argv:
            execute_new_odin()

        else:
            # if user has forgotten to provide env or --env in a legitimate command
            execute_new_odin()
    # If env or --env is present
    elif "env" in sys.argv or "--env" in sys.argv:
        if "set" in sys.argv and "env" in sys.argv:
            if "--name" in sys.argv:
                env_name = sys.argv[sys.argv.index("--name") + 1]
                custom_cmd = "set env " + env_name
                arg_list = shlex.split(custom_cmd)
                execute_new_odin_with_custom_cmd(arg_list)
                execute_old_odin()
                return
            else:
                print("name not provided in set env command")
                exit(0)

        if "list" in sys.argv and "env" in sys.argv:
            if "--help" in sys.argv:
                execute_new_odin()
            else:
                old_env_list = subprocess.check_output([OLD_ODIN] + sys.argv[1:])
                new_env_list = subprocess.check_output([NEW_ODIN] + sys.argv[1:])
                display_all_envs(old_env_list, new_env_list)
                return

        if "describe" in sys.argv and "env" in sys.argv:
            if "--env" in sys.argv:
                env_name = sys.argv[sys.argv.index("--env") + 1]
            elif "--name" in sys.argv:
                env_name = sys.argv[sys.argv.index("--name") + 1]

            # Check if env exists in old Odin first
            if check_env_exists_in_old_odin(env_name):
                execute_old_odin()
            else:
                execute_new_odin()

        service_name = None
        # env_name = None
        if "--env" in sys.argv:
            env_name = sys.argv[sys.argv.index("--env") + 1]
            if "--file" in sys.argv:
                file_index = sys.argv.index("--file") + 1
                file_path = sys.argv[file_index]
                service_name = get_service_name_from_file(file_path)
            elif "--service" in sys.argv:
                # need to check for --service strictly before --name for operate
                service_name = sys.argv[sys.argv.index("--service") + 1]
            elif "--name" in sys.argv:
                service_name = sys.argv[sys.argv.index("--name") + 1]
        elif "--name" in sys.argv:
            env_name = sys.argv[sys.argv.index("--name") + 1]
            if "--service" in sys.argv:
                service_name = sys.argv[sys.argv.index("--service") + 1]

        if env_name is not None and check_env_exists_in_old_odin(env_name):
            if service_name is not None and is_service_migrated_to_new_odin(service_name, env_name):
                execute_new_odin()
            else:
                execute_old_odin()
        else:
            execute_new_odin()
    else:
        execute_new_odin()


if __name__ == '__main__':
    update_binary()
    set_tokens(os.path.expanduser("~/.odin/config"))
    main()