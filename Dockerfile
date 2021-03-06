FROM golang
WORKDIR /toet
COPY . /toet
EXPOSE 11111
RUN ["mkdir", "/logs"]
RUN ["touch", "/logs/server.log"]
RUN ["go", "mod", "tidy"]
CMD ["go", "run", "main.go"]
#CMD ["tail", "-f", "main.go"]
