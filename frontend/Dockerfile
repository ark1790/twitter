# build stage
FROM node:10.15.3 as build-stage

RUN apt-get update && apt-get install -y jq python

RUN mkdir /app
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

ARG NODE_ENV=development
ENV NODE_ENV ${NODE_ENV}

ARG NETRC
ARG SSH_PRIVATE_KEY
RUN echo "${NETRC}" > /root/.netrc

# production stage
FROM nginx as production-stage

RUN rm /etc/nginx/conf.d/default.conf
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx-conf/default.conf /etc/nginx/conf.d/

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]