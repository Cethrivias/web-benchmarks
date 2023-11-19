FROM node:21 as base

WORKDIR /home/node/app

COPY src ./src
COPY package*.json ./
COPY tsconfig.json ./

RUN npm install
RUN npm run build
RUN npm prune --production

FROM node:21-slim
WORKDIR /home/node/app

COPY --from=base /home/node/app/dist ./dist
COPY --from=base /home/node/app/node_modules ./node_modules

ENTRYPOINT ["node", "dist"]
