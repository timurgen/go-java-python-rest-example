FROM iron/base

COPY whos-next-rest-go /opt/service/

WORKDIR /opt/service

RUN chmod +x /opt/service/whos-next-rest-go

EXPOSE 8080:8080

CMD /opt/service/whos-next-rest-go