# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Copy production environment file
COPY .env.production .env

# Build the application
RUN npm run build

# Expose the desired port (default for Next.js is 3000)
EXPOSE 3000

# Start the application
CMD ["npm", "run", "dev"]
