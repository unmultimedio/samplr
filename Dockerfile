FROM ubuntu:18.04
COPY . /app
#samplr#RUN make /mockapp
RUN make /secretapp
CMD python /app/app.py
