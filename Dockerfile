FROM aionnect/alpine-ffmpeg
WORKDIR /parsehtml
COPY . .
RUN chmod +x /parsehtml/main
EXPOSE 80
ENTRYPOINT ["./main"]
