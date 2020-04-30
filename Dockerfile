FROM aionnect/alpine-ffmpeg
WORKDIR /usr/src
COPY . . 
EXPOSE 81
ENTRYPOINT ["./main"]
