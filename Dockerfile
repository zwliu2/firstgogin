FROM centos
WORKDIR /usr/src
COPY . . 
EXPOSE 81
CMD ["./main"]
