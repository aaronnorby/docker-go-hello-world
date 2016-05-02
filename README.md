Create the app image (here called `hume`). Run this from the `app/` directory:

```
docker build -t hume .
```

Create your nginx image. Run this from the `/nginx` directory:

```
docker build -t hume-nginx .
```

Run it all (create the containers):
(each of these commands can be run from a separate terminal window or else pass the
`-d` flag to daemonize instead of the `--rm` flag (these are incompatible flags))

```
docker run --name hume1 -p 8000 --rm hume
docker run --name hume2 -p 8000 --rm hume
docker run --name hume3 -p 8000 --rm hume
docker run --name nginx -p 80:80 --link hume1:hume1 --link hume2:hume2 --link hume3:hume3 --rm hume-nginx
```

When we run `hume-nginx` all those links are linking the other containers to what
the nginx server is configured to look for (in this case, by those same names) in
`nginx.conf`.

This will make the application available on port 80 at the ip address of the host
machine. If you're using docker-machine (like on a Mac) then this will not be
`localhost`, because it's running a VirtualBox machine that the containers are
living in. To get the ip address run in your docker terminal:

```
docker-machine ip default
```

Go to `<machine_ip>/count` to get a count of the number of visits. You'll see that
the count jumps around a bit, showing the load balancer at work (each proxied
server keeps its own count of how many times its gotten requests).

Alternatively, you can make it all happen at once via docker-compose via the
`docker-compose.yml` file. From the directory where that file lives, run:

```
docker-compose up -d
```

The `-d` flag runs it in the background. To see running processes now, instead of
`docker ps` you run `docker-compose ps`. And to stop the process you run
`docker-compose stop`.
