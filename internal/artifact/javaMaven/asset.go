package javaMaven

func DockerFile() string {
	return `
ARG MAVEN_VERSION
ARG JAVA_VERSION

FROM maven:${MAVEN_VERSION}-jdk-${JAVA_VERSION}-slim

ARG PORT
ARG RUN_COMMAND

COPY ./ /app/
WORKDIR /app

EXPOSE ${PORT}

CMD ["bash", "-c", ${RUN_COMMAND}]
`
}