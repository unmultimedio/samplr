FROM ubuntu:18.04
COPY . /app
#samplr#RUN make /mockapp
RUN make /app
CMD python /app/app.py
