FROM node:16-alpine

ENV HOST=0.0.0.0 PORT=3000 NODE_ENV=production
WORKDIR /app

COPY . .
RUN npm ci --production

USER node
EXPOSE 3000

CMD ["npm", "start"]
