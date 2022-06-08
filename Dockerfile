FROM node:16 AS build

WORKDIR /work

COPY package.json .
COPY yarn.lock .
RUN yarn install --no-progress

COPY . .
RUN yarn build

FROM node:16

WORKDIR /work

RUN yarn global add serve

COPY package.json .
COPY yarn.lock .
RUN yarn install --production --no-progress

COPY --from=build /work/dist ./dist

ENTRYPOINT [ "serve", "-s", "/work/dist" ]
