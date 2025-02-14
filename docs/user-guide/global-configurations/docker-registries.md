# Container Registries

The global configuration helps you add your `Container Registry`. In the container registry, you provide credentials of your registry, where your images will be stored. And this will be shown to you as a drop-down on `Docker Build Config` Page.

## Add Container Registry configuration:

Go to the `Container Registry` section of `Global Configuration`. Click on `Add container registry`.

You will see below the input fields to configure the container registry.

* Name
* Registry type
  * ecr
    * AWS region
    * Access key ID
    * Secret access key
  * docker hub
    * Username
    * Password
  * Others
    * Username
    * password
* Registry URL
* Set as default

![](../../user-guide/global-configurations/images/Container_Registry.jpg)

### Name

Provide a name to your registry, this name will be shown to you in Docker Build Config as a drop-down.

### Registry type

Here you can select the type of the Registry. We are supporting three types- `docker hub`, `ecr` and `others`. You can select any one of them from the drop-down. By default, this value is `ecr`. If you select ecr then you have to provide some information like- `AWS region, Access Key`, and `Secret Key`. If you select docker hub then you have to provide `Username` and `Password`. And if you select others then you have to provide the `Username` and `Password`.

### Registry URL

Select any type of Registry from the drop-down, you have to provide the URL of your registry. Create your registry and provide the URL of that registry in the URL box.

### Registry Type- ECR:

You have to provide the below information if you select the registry type as ECR.

* **AWS region**

Select your AWS region from the drop-down, region where you have created your registry in.

* **Access key ID**

Inside the Access key ID box, provide your AWS access key.

* **Secret access key**

Provide your AWS secret access key ID.

![](../../user-guide/global-configurations/images/Container_Registry_ecr.jpg)

### Registry Type- Docker Hub 

You have to provide the below information if you select the registry type as Docker Hub.

* **Username**

Give the username of the docker hub account you used for creating your registry in.

* **Password**

Give the password/[token](https://docs.docker.com/docker-hub/access-tokens/) corresponding to your docker hub account.

![](../../user-guide/global-configurations/images/Container_Registry_DockerHub.jpg)

### Registry Type Others:

You have to provide the below information if you select the registry type as others.

* **Username**

Give the username of your account, where you have created your registry in.

* **Password**

Give the password corresponding to the username of your registry.

![](../../user-guide/global-configurations/images/Container_Registry_others.jpg)

### Set as default:

If you enable the `Set as default` option, then this registry name will be set as default in the `Container Registry` section inside the `Docker build config` page. This is optional. You can keep it disabled.

### Advance Registry Url connection options:

* If you enable the `Allow Only Secure Connection` option, then this registry allows only secure connections.

* If you enable the `Allow Secure Connection With CA Certificate` option, then you have to upload/provide private CA certificate (ca.crt).

* If the container registry is insecure (for eg : SSL certificate is expired), then you enable the `Allow Insecure Connection` option.

Now click on `Save` to save the configuration of the `Container registry`.

### Note:

You can use any registry which can be authenticated using `docker login -u <username> -p <password> <registry-url>`. However these registries might provide a more secured way for authentication, which we will support later.
Some popular registries which can be used using username and password mechanism:

* **Google Container Registry (GCR)** : JSON key file authentication method can be used to authenticate with username and password. Please follow [link](https://cloud.google.com/container-registry/docs/advanced-authentication#json-key) for getting username and password for this registry. Please remove all the white spaces from json key and wrap it in single quote while putting in password field.  

![](../../user-guide/global-configurations/images/Container_Registry_gcr.jpg)

* **Google Artifact Registry (GAR)** : JSON key file authentication method can be used to authenticate with username and password. Please follow [link](https://cloud.google.com/artifact-registry/docs/docker/authentication#json-key) for getting username and password for this registry. Please remove all the white spaces from json key and wrap it in single quote while putting in password field.
* **Azure Container Registry (ACR)** : Service principal authentication method can be used to authenticate with username and password. Please follow [link](https://docs.microsoft.com/en-us/azure/container-registry/container-registry-auth-service-principal) for getting username and password for this registry.

## Integrating With External Container Registry

If you want to use a private registry for container registry other than ecr, this will be used to push image and then create a secret in same environment to pull the image to deploy. To create secret, go to charts section and search for chart ‘dt-secrets’ and configure the chart. Provide an App Name and select the Project and Environment in which you want to deploy this chart and then configure the values.yaml as shown in example. The given example is for DockerHub but you can configure similarly for any container registry that you want to use.

```yaml
name: regcred
type: kubernetes.io/dockerconfigjson
labels:
 test: chart
secrets:
 data:
   - key: .dockerconfigjson
     value: '{"auths":{"https://index.docker.io/v1/":{"username":"<username>","password":"<password>}}}'
```     

The `name` that you provide in values.yaml ie. `regcred` is name of the secret that will be used as `imagePullSecrets` to pull the image from docker hub to deploy. To know how `imagePullSecrets` will be used in the deployment-template, please follow the [documentation](https://docs.devtron.ai/devtron/user-guide/creating-application/deployment-template/rollout-deployment#imagepullsecrets).
