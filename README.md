# member-club

The application consists of two components: REST API, which is implemented in Golang and UI front, which is implemented in React.JS.<br/>
To simplify the task of deployment and development, a monorep approach was chosen.

## Table of contents

+ [API](#API);
+ [Web](#Web);
+ [How to run](#Run);
<br/>
<br/>

## <a name="API"></a> API

Golang REST API service.<br/>
Description of REST endpoints are described in the [swagger.yaml](./api/swagger.yaml) document.<br/>
For implementation, the Gin framework was selected and the generation of server files based on the swagger document and utility [oapi-codegen](https://github.com/deepmap/oapi-codegen)

## <a name="Web"></a> Web

React.JS application.<br/>
Implements a single page that uses the endpoints API to add a club member and display a list of all members.<br/>
For additional info read [README.md](./web/README.md)

## <a name="Run"></a> How to run

For a successful launch, you must complete the following steps:

1. To start the application, you need to run the command in the root folder of the project:
    > ```
    > docker-compose up -d
    > ```
2. After successfully completing all the steps, the following result should appear in the console:
    > ```
    > [+] Running 3/3
    >  ⠿ Network member-club_default  Created
    >  ⠿ Container member-club_web_1  Started
    >  ⠿ Container member-club_api_1  Started
    > ```
3. This says that the applications have been successfully launched. Next, you need to go to the browser at:
    > ```
    > http://localhost:3000/
    > ```