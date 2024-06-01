# Telegram Cloud

Telegram Cloud is a software that allows you to host files up to 50MB on Telegram. This documentation provides a comprehensive guide on how to install, configure, and use the Telegram Cloud software.

## Table of Contents

1. [Installation](#installation)
2. [Configuration](#configuration)
3. [Usage](#usage)
4. [API Endpoints](#api-endpoints)
5. [Safe Usage Policy](#safe-usage-policy)

## Installation

To install Telegram Cloud, follow these steps:

1. **Clone the Repository**: Clone the Git repository to your local machine.

    ```sh
    git clone https://github.com/yourusername/telegram-cloud.git
    cd telegram-cloud
    ```

2. **Set Up Environment Variables**: Change the environment variables as required. You need to set up the Telegram bot token and your public or private channel username.

    - **Bot Token**: Obtain a bot token from [BotFather](https://t.me/botfather). Start a conversation with BotFather, use the `/newbot` command to create a new bot, and get the token.

    ![BotFather Token](docs/botfather.png)

    - **Channel Username**: Obtain your public or private channel username. To find your channel username, open your Telegram app, go to the channel, and note the username.

    ![Channel Username](docs/channel-username.png)

    **Note**:  Add Bot to the channel and grant it all permissions

    - **Environment Variables**: Create a `.env` file in the project root directory and add the following:

    ```env
    TELEGRAM_BOT_TOKEN=your_bot_token_here
    TELEGRAM_CHANNEL_USERNAME=your_channel_username_here
    ```

3. **Build the Application**: Use the Go toolchain to build the application.

    ```sh
    go build -o telegram-cloud
    ```

4. **Run the Application**: Start the application by running the binary.

    ```sh
    ./telegram-cloud
    ```

## Configuration

To configure Telegram Cloud, ensure you have set up the following environment variables in your `.env` file:

- `TELEGRAM_BOT_TOKEN`: Your Telegram bot token obtained from BotFather.
- `TELEGRAM_CHANNEL_USERNAME`: Your public or private channel username.

Example `.env` file:

```env
TELEGRAM_BOT_TOKEN=123456789:ABCdefGHIjklMNOpqrSTUvwxYZ
TELEGRAM_CHANNEL_USERNAME=@my_channel
```

## Usage

After setting up and running the application, you can upload files to Telegram by making a POST request to the `/upload` endpoint.

### Check the Application

To ensure the application is running correctly, access the following endpoint:

```sh
http://localhost:8080/upload
```
## API Endpoints

### Upload File

**Endpoint**: `POST /upload`

**Description**: Uploads a file to Telegram. Accepts a single file and returns a response including the direct URL to the file.

**Request**:

- **Method**: `POST`
- **URL**: `http://localhost:8080/upload`
- **Content-Type**: `multipart/form-data`
- **Parameters**:
  - `file` (required): The file to upload.

**Response**:

- **Success**: Returns the direct URL to the uploaded file.
- **Error**: Returns an error message if the file upload fails.

**Example**:

```sh
curl -X POST -F "file=@/path/to/your/file" http://localhost:8080/upload
```
## Safe Usage Policy

When using Telegram Cloud, ensure you follow these safety guidelines:

1. **File Size**: Only upload files that are 50MB or smaller. Will include support for larger files in the future

2. **File Types**: Ensure the files you upload do not violate Telegram's terms of service. Avoid uploading malicious or illegal content.

3. **Private Data**: Do not upload sensitive or confidential data without proper encryption and security measures.

4. **API Usage**: Do not abuse the API by uploading a large number of files in a short period. Implement rate limiting if necessary.

5. **Data Privacy**: Be aware that files uploaded to Telegram may be accessible to others if not properly secured. Use private channels for sensitive data.

## Conclusion

Telegram Cloud provides a convenient way to host files on Telegram. By following this documentation, you can easily install, configure, and use the software to upload and manage your files. Ensure you adhere to the safe usage policy to maintain the integrity and security of your data.

For any issues or contributions, feel free to open an issue or submit a pull request on the [GitHub repository](https://github.com/charlesozo/telegram-cloud-storage).


