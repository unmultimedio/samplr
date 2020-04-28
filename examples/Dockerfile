#hsamplr#FROM approved:image
FROM my:test:image

#hsamplr#RUN echo "some command"
RUN echo "some command" --with=flags --that-i-dont=want_to_share

COPY . /app

#ssamplr# This block of comments will only
#ssamplr# appear in the original file since
#ssamplr# it is preceded by the
#ssamplr# secret #ssamplr# keyword

#samplr#RUN make /publicapp
RUN make /privateapp

#samplr#CMD some_command /publicapp/foo
CMD some_command /privateapp/foo
