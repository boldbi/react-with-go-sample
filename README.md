# Bold BI Embedded Sample in React with Go

This Bold BI React with Go sample contains the Dashboard embedding sample. In this sample React application act as the front-end and the Go sample act as the back-end application. This application aims to demonstrate how to render the dashboard available on your Bold BI server.

## Dashboard view

![Dashboard View](/images/dashboard.png)

## Requirements/Prerequisites

* [Go installer](https://go.dev/dl/)
* [Visual Studio Code](https://code.visualstudio.com/download)
* [Node.js](https://nodejs.org/en/)

 > **NOTE:** Node.js v18.17 to v20.15 are supported.

### Supported browsers
  
* Google Chrome, Microsoft Edge, and Mozilla Firefox.

## Configuration

* Please ensure you have enabled embed authentication on the `embed settings` page. If it is not currently enabled, please refer to the following image or detailed [instructions](https://help.boldbi.com/site-administration/embed-settings/#get-embed-secret-code?utm_source=github&utm_medium=backlinks) to enable it.

    ![Embed Settings](/images/enable-embedsecretkey.png)

* To download the `embedConfig.json` file, please follow this [link](https://help.boldbi.com/site-administration/embed-settings/#get-embed-configuration-file?utm_source=github&utm_medium=backlinks) for reference. Additionally, you can refer to the following image for visual guidance.

     ![Embed Settings Download](/images/download-embedsecretkey.png)
     ![EmbedConfig Properties](/images/embedconfig-file.png)

* Copy the downloaded `embedConfig.json` file and paste it into the designated [location](https://github.com/boldbi/react-with-go-sample/tree/master/Go) within the application. Please ensure you have placed it in the application, as shown in the following image.

    ![EmbedConfig image](/images/embedconfig-location.png)

## Run a Sample Using Command Line Interface

* Open the command line interface and navigate to the specified file [location](https://github.com/boldbi/react-with-go-sample/tree/master/Go) where the project is located.

* Run the back-end `Go` sample using the following command `go run main.go`.

* Open the command line interface and navigate to the specified file [location](https://github.com/boldbi/react-with-go-sample/tree/master/React) where the project is located.

* To install all dependent packages, use the following command `npm install`.

* Finally, run the application using the command `npm start`. After executing the command, the application will automatically launch in the default browser. You can access it at the specified port number (e.g., <http://localhost:3000/>).

## Developer IDE

* [Visual Studio Code](<https://code.visualstudio.com/download>)

### Run a Sample Using Visual Studio Code

* Open the `Go` sample in Visual Studio Code.

* Install the extension `Go` in Visual Studio Code. Please refer to the following image.
    ![Extension](/images/go-extension.png)

* Run the back-end `Go` sample using the following command `go run main.go`.

* Open the `React` sample in a new window of Visual Studio Code.

* To install all dependent packages, use the following command `npm install`.

* Finally, run the application using the command `npm start`. After executing the command, the application will automatically launch in the default browser. You can access it at the specified port number (e.g., <http://localhost:3000/>).

![dashboard image](/images/dashboard.png)

Please refer to the [help documentation](https://help.boldbi.com/embedding-options/embedding-sdk/samples/react-with-go/#how-to-run-the-sample?utm_source=github&utm_medium=backlinks) to know how to run the sample.

## Online Demos

Look at the Bold BI Embedding sample to live demo [here](https://samples.boldbi.com/embed?utm_source=github&utm_medium=backlinks).

## Documentation

 A complete Bold BI Embedding documentation can be found on [Bold BI Embedding Help](https://help.boldbi.com/embedded-bi/javascript-based/?utm_source=github&utm_medium=backlinks).
