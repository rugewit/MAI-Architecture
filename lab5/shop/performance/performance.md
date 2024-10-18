<h2>Используя Redis</h2>
<pre>(base) <font color="#26A269"><b>rugewit@rugewit</b></font>:<font color="#12488B"><b>~</b></font>$ wrk -t12 -c400 -d30s http://0.0.0.0:8081/users/665662d9ea106dc5807716bc
Running 30s test @ http://0.0.0.0:8081/users/665662d9ea106dc5807716bc
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.92ms    3.13ms  66.02ms   73.29%
    Req/Sec     4.18k   420.12    12.68k    83.54%
  1497426 requests in 30.10s, 501.25MB read
Requests/sec:  49750.66
Transfer/sec:     16.65MB
</pre>
<h2>Не используя Redis</h2>
<pre>
(base) <font color="#26A269"><b>rugewit@rugewit</b></font>:<font color="#12488B"><b>~</b></font>$ wrk -t12 -c400 -d30s http://0.0.0.0:8081/users/665662d9ea106dc5807716bc
Running 30s test @ http://0.0.0.0:8081/users/665662d9ea106dc5807716bc
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    16.46ms    4.54ms  56.17ms   72.56%
    Req/Sec     2.01k   143.71     5.23k    83.51%
  721409 requests in 30.10s, 241.48MB read
Requests/sec:  23969.49
Transfer/sec:      8.02MB
</pre>