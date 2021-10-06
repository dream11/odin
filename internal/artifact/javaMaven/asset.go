package javaMaven

func DockerFile() string {
	return `
ARG MAVEN_VERSION
ARG JAVA_VERSION

FROM maven:${MAVEN_VERSION}-jdk-${JAVA_VERSION}-slim

ARG JAR_PATH
ARG PORT
ARG RUN_COMMAND

COPY ${JAR_PATH} /app/
WORKDIR /app

COPY --chown=app:app /app/${JAR_PATH} /app/${JAR_NAME}

USER app
EXPOSE ${PORT}

CMD ["bash", "-c", ${RUN_COMMAND}]
`
}