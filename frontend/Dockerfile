# Start with the base Node.js image
FROM node:20.2.0-bullseye-slim

# Set the working directory
WORKDIR /app

# # Copy the package.json and yarn.lock files
# COPY package.json ./
# COPY yarn.lock ./

# Update package lists and install Python 3 and other required packages
RUN apt-get update && apt-get install -y python3 python3-pip

# Optionally, set up a symbolic link to make Python 3 the default
RUN ln -s /usr/bin/python3 /usr/bin/python

# Copy the rest of the application code
COPY . .

# Install the dependencies
RUN yarn install

# Check svelte is working
RUN yarn check

# Build the Svelte application
RUN yarn build

# Expose the port on which the application will run
EXPOSE 80

# Set the entrypoint for the Docker container
CMD ["yarn", "start"]