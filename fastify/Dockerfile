FROM node:21

WORKDIR /home/node/app

COPY src ./src
COPY package*.json ./
COPY tsconfig.json ./

RUN npm install

RUN npm run build

EXPOSE 3000

ENTRYPOINT ["node", "dist"]
