# [Ngrok](https://ngrok.com/)

I now use [Cloudflare Tunnel](https://twitter.com/signalnerve/status/1449038210076655624) for exposing local services to the web (such as localhost).

## Using it

If I started my local server on port 3000. This command: `ngrok http 3000` will create a shareable link of my tunneled server.

## Notes

- [ngrok lets you inspect incoming requests in a web UI, and even replay past requests.](https://twitter.com/geoffreylitt/status/1379092674280579082)

## Links

- [Node wrapper for ngrok](https://github.com/bubenshchykov/ngrok)
- [ngrok code](https://github.com/inconshreveable/ngrok)
- [pgrok](https://github.com/jerson/pgrok) - Free Introspected tunnels to localhost, like ngrok but free and unlimited.
- [Setting up Cloudflare Tunnel for development (2021)](https://kirillplatonov.com/posts/setting-up-cloudflare-tunnel-for-development/)
- [JPRQ](https://github.com/azimjohn/jprq) - Get Your Localhost Online and HTTPS. Ngrok Alternative.
- [Roll your own Ngrok with Nginx, Letsencrypt, and SSH reverse tunnelling (2019)](https://jerrington.me/posts/2019-01-29-self-hosted-ngrok.html) ([HN](https://news.ycombinator.com/item?id=30891494))
