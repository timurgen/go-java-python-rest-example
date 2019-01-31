FROM openjdk:8-jre-alpine
COPY ./target/whosnext-0.0.1-SNAPSHOT.jar /opt/service/whosnext-0.0.1-SNAPSHOT.jar
ENTRYPOINT ["/usr/bin/java"]
CMD ["-jar", "/opt/service/whosnext-0.0.1-SNAPSHOT.jar"]
EXPOSE 8080:8080