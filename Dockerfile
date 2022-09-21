FROM golang:1.18
ENV PORT 8000
#
#RUN apt-get update
#RUN apt-get install make

RUN mkdir -p /usr/src/codekata/backend/
WORKDIR /usr/src/codekata/backend

COPY . /usr/src/codekata/backend
EXPOSE 8000

CMD "make" "build"
CMD "make" "run"
