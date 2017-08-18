FROM alpine:latest
ENV PORT=8080
RUN mkdir /app
COPY echo-server-linux /app/echo-server
EXPOSE 8080
CMD /app/echo-server --port=$PORT

