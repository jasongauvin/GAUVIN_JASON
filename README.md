# Partial of Software Architecture and Back Development. 
### Github link : https://github.com/jasongauvin/GAUVIN_JASON

This is based on my template repository "gogoleplate" made for golang.
-   hot reload with air 
-   postgres database
-   github configuration
-   uses conventions
-   JWT
-   Authorization API

It is available at the following address : https://github.com/jasongauvin/gogoleplate

This application must make with an API, to manage questionnaires with the possibility of resuming lessons and saving their progress. User and administrator accounts must be usable with authentication and authorization. 

## Todo

- [x] Docker
- [x] Documentation
- [x] JWT
- [x] Autorisation
- [x] Saving progress
- [x] Customers controller
- [x] Forms controller
- [x] Questions controller
- [x] Types Controller
- [ ] Answers controller
- [ ] Unit test
## Requirements

To run this project you'll need to have docker and yarn installed on your machine.
If you want to develop on this project it's recommended to have `golang` installed on your machine.

## Project setup

First you need to created your `.env` file (you can use the .env.dist file).
You will placed at the root of your project : `public.pem` and `private.pem`


```sh
  docker-compose up --build
```

### RSA keys (needed to generate and read tokens)

You will need 2 files placed at the root of your project : `public.pem` and `private.pem`.

```sh
    # use the following password: private_key
    openssl genrsa -des3 -out private.pem 2048
    openssl rsa -in private.pem -outform PEM -pubout -out public.pem
```
# Commit naming conventions

If you want to contribute to the project you'll need to name your commits according to the following convention :

    type: description.

The following types are available :

-   **feat** : Feature commit.
-   **fix** : Fixing a bug.
-   **update** : Updating code or dependencies.
-   **revert** : Removing changes.
-   **doc** : Add changes to the documentation.
-   **refacto** : Refactoring code.
-   **build** : Modifications linked to the infrastructure.
-   **param** : Modifications on the config.

## Branch naming convention

The branch should have a name that reflects it's purpose.

The convention is to prefix the branch name with `feature-`. All the words must be separated by a `-`.

Branch name example : `feature-user-authentication`