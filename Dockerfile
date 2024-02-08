# Start from a base Node.js image
FROM node:latest

# Set working directory
WORKDIR /usr/src/app

# Copy package.json and package-lock.json (if available)
COPY package*.json ./

# Install dependencies, including bcrypt
RUN npm install

# Copy the rest of the application code
COPY . .

# Expose port if needed
# EXPOSE 3000

# Command to run the application
CMD ["node", "app.js"]
