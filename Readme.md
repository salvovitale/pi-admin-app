# Pi Admin App

A simple web application built using Go that allows you to manage some functionalities of your Raspberry Pi without the need of a terminal. The application serves an HTML page with buttons to perform various actions, such as rebooting the Raspberry Pi.

## Features

- Web-based interface
- Reboot Raspberry Pi with a button click
- Easy to extend with more buttons and functionalities
- Serves static files for JavaScript and other assets
- Can be run as a daemon service on Raspberry Pi

## Project Structure
This Go application is organized to keep HTML templates, JavaScript functions, and server-side handlers separate. You can easily add more buttons and functionalities by following these steps:

1. Add a new button with an onclick event handler in the HTML template in the serveHTML function.
2. Define a new JavaScript function for the button's onclick event in the main.js file.
3.  Create a new server-side handler function in the main.go file, if required.
4. Register the new handler function with a URL path in the main() function.


## Running the App as a Daemon Service

To make the Go app run automatically when the Raspberry Pi restarts, you can create a systemd service. Here are the steps:

1. Build your Go application, if you haven't done so already:

```bash
go build main.go
```

2. Move the compiled binary to a suitable location, like /usr/local/bin:

```bash
sudo mv main /usr/local/bin/pi-admin-app
```

3. Create a systemd service file for your application by creating a new file /etc/systemd/system/pi-admin-app.service:

```bash
sudo nano /etc/systemd/system/pi-admin-app.service
```

4. Add the following contents to the service file:

```bash
[Unit]
Description=Raspberry Pi Control Panel Web Server
After=network.target

[Service]
User=root
Group=root
ExecStart=/usr/local/bin/pi-admin-app
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

This configuration specifies that the service should start after the network target is reached, run as root, and restart if the process fails. Adjust the **User** and **Group** fields if you want to run the service under a different user.

5. Reload the systemd configuration:

```bash
sudo systemctl daemon-reload
```

6. Enable the service to start on boot:

```bash
sudo systemctl enable pi-admin-app.service
```

7. Start the service:

```bash
sudo systemctl start pi-admin-app.service
```

8. Check the status of the service:

```bash
sudo systemctl status pi-admin-app.service
```