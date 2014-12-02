## extremely naive Docker Hub webhook tester

run it using `make`, then add `http://<hostip>/whatever` to a webhook

I get the following output:


```
docker run --rm -it -p 80:80 hook
+ exec app


 push_data
   pushed_at is float64 1.417564071e+09
   images is an array:
   0 imagehash1
   1 imagehash2
   2 imagehash3
   pusher is string svendowideit

 callback_url is string https://registry.hub.docker.com/u/svendowideit/testhook/hook/2141abe0dd12gabebh11i4d0e30242eg110012/
 repository
   namespace is string svendowideit
   star_count is float64 0
   comment_count is float64 0
   date_created is float64 1.417494799e+09
   description is string 
   is_official is bool false
   name is string testhook
   full_description is nil
   repo_url is string https://registry.hub.docker.com/u/svendowideit/testhook/
   is_private is bool true
   repo_name is string svendowideit/testhook
   dockerfile is string #
# BUILD         docker build -t svendowideit/apt-cacher .
# RUN           docker run -d -p 3142:3142 -name apt-cacher-run apt-cacher
#
# and then you can run containers with:
#               docker run -t -i -rm -e http_proxy http://192.168.1.2:3142/ debian bash
#
FROM            ubuntu
MAINTAINER      SvenDowideit@home.org.au


VOLUME          ["/var/cache/apt-cacher-ng"]
RUN             apt-get update ; apt-get install -yq apt-cacher-ng

EXPOSE          3142
CMD             chmod 777 /var/cache/apt-cacher-ng ; /etc/init.d/apt-cacher-ng start ; tail -f /var/log/apt-cacher-ng/*

   status is string Active
   is_trusted is bool true
   owner is string svendowideit

callback to  https://registry.hub.docker.com/u/svendowideit/testhook/hook/2141abe0dd12gabebh11i4d0e30242eg110012/
"OK"
```

once you know the callback_url, you can just as easily use curl :)

```
$ curl --data '{"state":"success"}' https://registry.hub.docker.com/u/svendowideit/testhook/hook/21402b11bee3ecb1i4fifg0242eg1100a1/
"OK"
```
