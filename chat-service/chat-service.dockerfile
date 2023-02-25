# build a tiny docker image
FROM alpine

# create folder in root
RUN mkdir /app
RUN mkdir /app/download

# copy exe from previous image to new image
COPY chatServiceApp /app

# provide defaults for an executing container
CMD [ "/app/chatServiceApp" ]