FROM node:18.12.1-alpine

WORKDIR /front

COPY package.json .
COPY yarn.lock .

RUN yarn && yarn cache clean

COPY . .

EXPOSE 3000

CMD ["yarn", "run", "dev"]
