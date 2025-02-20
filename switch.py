import glob
import os
import re
import ssl
import subprocess
import sys
import json
from collections import defaultdict

try:
    from urllib.request import Request, urlopen, urlretrieve  # Python 3
except ImportError:
    from urllib2 import Request, urlopen, urlretrieve  # Python 2

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


def update_binary():
    url = f"https://api.github.com/repos/dream11/odin/git/trees/{branch}?recursive=1"

    req = Request(url)
    req.add_header('Accept', 'application/json')
    try:
        response = urlopen(req).read()
        if response:
            data = json.loads(response)
            files = [item["path"] for item in data.get("tree", []) if item["type"] == "blob"]

            # Filter files that start with "odin"
            filtered_files = [file for file in files if file.startswith('odin-')]
            latest_version = filtered_files[0].split("-")[1]

            current_version = get_current_bin_version()

            if current_version is None or current_version < latest_version:
                print(f"Updating odin binary to version {latest_version}")
                url = f"https://raw.githubusercontent.com/dream11/odin/{branch}/odin-{latest_version}"
                filepath = os.path.join(INSTALL_DIR, f"odin-{latest_version}")
                urlretrieve(url, filename=filepath)
                os.chmod(filepath, 0o755)
                # Enable app verification and remove quarantine attributes
                subprocess.call(["sudo", "spctl", "--master-enable"])
                subprocess.call(["xattr", "-dr", "com.apple.quarantine", filepath])
        else:
            print(f"Error: Unable to fetch files (Status Code: {response.status_code})")
            return []
    except Exception as e:
        print(f"Error: Unable to fetch files: {e}")


def execute_new_odin():
    subprocess.call([NEW_ODIN] + sys.argv[1:])
    exit(0)


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


def check_env_exists(env_name):
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


def check_service_migrated(service_name, env_name):
    global odin_access_token, odin_backend_address, checkMigrationStatusUri
    url = odin_backend_address + checkMigrationStatusUri + "/" + env_name + "/" + service_name
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
            name, team, created_by, env_type, state, account = env_data[0], env_data[1], env_data[2], env_data[3], env_data[4], env_data[5]
            envs.append({
                "name":  name.strip(),
                "team": team.strip(),
                "created_by": created_by.strip(),
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
    grouped_envs = defaultdict(lambda: {"team": "", "created_by": "", "env_type": "", "state": "", "accounts": set()})

    for env in envs:
        key = env["name"].strip() or "N/A"
        grouped_envs[key]["team"] = env.get("team", "").strip() or "N/A"
        grouped_envs[key]["created_by"] = env.get("created_by", "").strip() or "N/A"
        grouped_envs[key]["env_type"] = env.get("env_type", "").strip() or "N/A"
        grouped_envs[key]["state"] = env.get("state", "").strip() or "N/A"
        if env["account"].strip():
            grouped_envs[key]["accounts"].add(env["account"].strip())  # Use set to remove duplicates

    # Convert accounts set to a comma-separated string (or "N/A" if empty)
    for key in grouped_envs:
        grouped_envs[key]["accounts"] = ", ".join(sorted(grouped_envs[key]["accounts"])) if grouped_envs[key]["accounts"] else "N/A"

    # Define column headers
    headers = ["NAME", "TEAM", "CREATED BY", "ENV TYPE", "STATE", "ACCOUNTS"]

    # Determine column widths dynamically
    col_widths = [
        max(len(key) for key in list(grouped_envs.keys()) + [headers[0]]),
        max(len(row["team"]) for row in list(grouped_envs.values()) + [{"team": headers[1]}]),
        max(len(row["created_by"]) for row in list(grouped_envs.values()) + [{"created_by": headers[2]}]),
        max(len(row["env_type"]) for row in list(grouped_envs.values()) + [{"env_type": headers[3]}]),
        max(len(row["state"]) for row in list(grouped_envs.values()) + [{"state": headers[4]}]),
        max(len(row["accounts"]) for row in list(grouped_envs.values()) + [{"accounts": headers[5]}]),
    ]

    # Print the header row
    header_format = " | ".join(f"{{:<{w}}}" for w in col_widths)
    print(header_format.format(*headers))
    print("-" * (sum(col_widths) + len(col_widths) * 3 - 3))  # Print separator

    # Print each grouped row
    for name, details in grouped_envs.items():
        print(header_format.format(name, details["team"], details["created_by"], details["env_type"], details["state"], details["accounts"]))


def main():
    global odin_access_key, odin_secret_access_key, odin_access_token, odin_backend_address, OLD_ODIN
    # If /opt/homebrew/bin/odin exists, use this as old_odin
    if os.path.isfile("/opt/homebrew/bin/odin"):
        OLD_ODIN = "/opt/homebrew/bin/odin"

    if len(sys.argv) == 1:
        print("No arguments provided")
        sys.exit(1)

    if len(sys.argv) == 2 and "--version" in sys.argv:
        execute_old_odin()

    if len(sys.argv) == 2 and "version" in sys.argv:
        execute_new_odin()

    if "configure" in sys.argv:
        config_file = os.path.expanduser("~/.odin/config")

        # Delete ~/.odin/config
        if os.path.isfile(config_file):
            os.remove(config_file)

        # Call configure for both old and new odin
        configure()

        exit(0)

    if is_dreampay():
        execute_old_odin()

    if "env" not in sys.argv and "--env" not in sys.argv:
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
    # If env or --env is present
    else:
        if "list" in sys.argv and "env" in sys.argv:
            old_env_list = subprocess.check_output([OLD_ODIN] + sys.argv[1:])
            new_env_list = subprocess.check_output([NEW_ODIN] + sys.argv[1:])
            display_all_envs(old_env_list, new_env_list)
            return

        if "--env" in sys.argv:
            env_name = sys.argv[sys.argv.index("--env") + 1]
            service_name = sys.argv[sys.argv.index("--name") + 1]
        else:
            env_name = sys.argv[sys.argv.index("--name") + 1]
            if "--service" in sys.argv:
                service_name = sys.argv[sys.argv.index("--service") + 1]

        if check_env_exists(env_name):
            if check_service_migrated(service_name, env_name):
                execute_new_odin()
            else:
                execute_old_odin()
        else:
            execute_new_odin()


if __name__ == '__main__':
    update_binary()
    set_tokens(os.path.expanduser("~/.odin/config"))
    main()
