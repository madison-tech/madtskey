## madtskey
Use this module to generate an ephemeral, one time use Auth key for your tailscale devices.

Once generated you can use this key to have new devices join your tailnet. We use this module as part of a larger Pulumi program when we want to have a cloud server provisioned by cloud-init and then join our tailnet.

To use it, first generate an [OAuth Client](https://login.tailscale.com/admin/settings/oauth) from your tailscale admin settings page.

Then create a .env file and copy the client ID, client Secret and [Organiation name](https://login.tailscale.com/admin/settings/general) as shown below. *Generally the first item in the settings page, if you are an individual account, the org name is usually your email address*:

```
OAUTH_CLIENT_ID=k123456CNTRL
OAUTH_CLIENT_SECRET=tskey-client-123456CNTRL-abcdefghijklmnopqrstuvwxyz
TAILNET=your_org_name
```

When using this automated approach, it is mandatory to set tags for these keys. You can add a tag by visiting the [Access Controls](https://login.tailscale.com/admin/acls/file) page and adding a tag to the root of the ACL e.g:

```
// Define the tags which can be applied to devices and by which users.
	"tagOwners": {
		"tag:mad-ts-key": ["autogroup:admin"],
	},
```

Then you can call `CreateAPIKey()` to create your API key. You have to specify how long you want this key to be alive for in seconds, a description for the key and the tags that you created in the ACL page:

```
key, err := CreateAPIKey(300, "my test key", []string{"tag:mad-ts-key"})
```
*We recommend keeping your expiry time very short since you will want to use this in an automation framework*

Then call `key.Key` to get your key which you can use when having new devices join your tailnet:

```
tailscale up --authkey=tskey-auth-k123456CNTRL-abcdefghijklmnopqrstuvwxyz
```
