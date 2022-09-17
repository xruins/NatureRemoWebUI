FROM node:carbon

WORKDIR /usr/src/nature-remo-web-ui

COPY package*.json .
RUN npm install
COPY . .

EXPOSE 3000

CMD ["npm", "start"]
