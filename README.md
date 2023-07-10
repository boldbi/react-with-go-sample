# Bold BI Embedded Sample in React with Go

This project was created using React with Go. In this sample React application act as the front-end and the Go sample act as the back-end application. This application aims to demonstrate how to render the dashboard available on your Bold BI server.

## Dashboard view

![Dashboard View](https://github.com/boldbi/aspnet-core-sample/assets/91586758/4af68f49-ffc0-400a-a323-55a3f3600a1d)

 ## Requirements/Prerequisites

 * [Go installer](https://go.dev/dl/)
 * [Visual Studio Code](https://code.visualstudio.com/download)
 * [Node.js](https://nodejs.org/en/)

 #### Help link

 * https://help.boldbi.com/embedded-bi/faq/where-can-i-find-the-product-version/

 #### Supported browsers
  
  * Google Chrome, Microsoft Edge, Mozilla Firefox.

 ## Configuration

  * Please ensure you have enabled embed authentication on the `embed settings` page. If it is not currently enabled, please refer to the following image or detailed [instructions](https://help.boldbi.com/site-administration/embed-settings/#get-embed-secret-code) to enable it.

    ![Embed Settings](https://github.com/boldbi/aspnet-core-sample/assets/91586758/b3a81978-9eb4-42b2-92bb-d1e2735ab007)

  * To download the `embedConfig.json` file, please follow this [link](https://help.boldbi.com/site-administration/embed-settings/#get-embed-configuration-file) for reference. Additionally, you can refer to the following image for visual guidance.

     ![Embed Settings Download](https://github.com/boldbi/aspnet-core-sample/assets/91586758/d27d4cfc-6a3e-4c34-975e-f5f22dea6172)
     ![EmbedConfig Properties](https://github.com/boldbi/aspnet-core-sample/assets/91586758/d6ce925a-0d4c-45d2-817e-24d6d59e0d63)

  * Copy the downloaded `embedConfig.json` file and paste it into the designated [location](https://github.com/boldbi/react-with-go-sample/tree/master/Go) within the application. Please ensure you have placed it in the application, as shown in the following image.
    
    ![EmbedConfig image](https://github.com/boldbi/aspnet-core-sample/assets/91586758/4d1489c0-ea7a-40ab-b067-e116ad9bee5a)

 ## Run a Sample Using Command Line Interface 
    
  * Open the command line interface and navigate to the specified file [location](https://github.com/boldbi/react-with-go-sample/tree/master/Go) where the project is located.

  * Run the back-end `Go` sample using the following command `go run main.go`.

  * Open the command line interface and navigate to the specified file [location](https://github.com/boldbi/react-with-go-sample/tree/master/React) where the project is located.
   
  * To install all dependent packages, use the following command `npm install`.

  * Finally, run the application using the command `npm start`. After executing the command, the application will automatically launch in the default browser. You can access it at the specified port number (e.g., http://localhost:3000/).

 ## Developer IDE

  * Visual studio code(https://code.visualstudio.com/download)

  ### Run a Sample Using Visual Studio Code
 
  * Open the `Go` sample in Visual Studio Code.

  * Install the extension `Go` in Visual Studio Code. Please refer to the following image.
    ![Extension](https://github.com/boldbi/aspnet-core-sample/assets/91586758/8cc5ca2f-f59f-4bd1-bb5c-3dc00ac1b2a8)

  * Run the back-end `Go` sample using the following command `go run main.go`.

  * Open the `React` sample in a new window of Visual Studio Code.
   
  * To install all dependent packages, use the following command `npm install`.

  * Finally, run the application using the command `npm start`. After executing the command, the application will automatically launch in the default browser. You can access it at the specified port number (e.g., http://localhost:3000/).

    ![dashboard image](https://github.com/boldbi/aspnet-core-sample/assets/91586758/4af68f49-ffc0-400a-a323-55a3f3600a1d)

Please refer to the [help documentation](https://help.boldbi.com/embedding-options/embedding-sdk/samples/react-with-go/#how-to-run-the-sample) to know how to run the sample.

## Online Demos

Look at the Bold BI Embedding sample to live demo [here](https://samples.boldbi.com/embed).

## Documentation

 A complete Bold BI Embedding documentation can be found on [Bold BI Embedding Help](https://help.boldbi.com/embedded-bi/javascript-based/).


