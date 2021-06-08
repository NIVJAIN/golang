 docker run -d --name mongodb \
      -e MONGO_INITDB_ROOT_USERNAME=root \
      -e MONGO_INITDB_ROOT_PASSWORD=iloveblockchain \
      -p 27017:27017 \
      mongo
