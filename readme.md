<style>
    body {
        padding: 0;
        margin: 0;
    }
    .title {
        font-size: 24px;
        display: flex;
        justify-content: center;
    }

</style>

<h1 class='title'>ptree</h1>
<span>trying to build a simple process memory attribution tool before dawn. lets see how it goes.</span>

<h3>progress</h3>
**should be changing as i proceed
<pre>
parent 0 has child(ren):
  └── 1 init
parent 479 has child(ren):
  └── 2741 go
parent 1 has child(ren):
  └── 276 syslogd
  └── 304 crond
  └── 403 udhcpc
  └── 467 getty
  └── 59 orbstack-agent:
parent 2741 has child(ren):
  └── 2845 main
parent 59 has child(ren):
  └── 479 sh
</pre>

<h3>testing</h3>
<em>disclaimer:</em> i am running this on my homelab(alpine 3.22, arm64 on orbstack vm). this should work on any linux distro.

<h3>requirements</h3>
<ul>
<li>any linux distro</li>
<li>>=go1.21</li>
</ul>

<h3>resource</h3>
<ul>
    <li>
        <a href="https://broman.dev/download/The%20Linux%20Programming%20Interface.pdf">The Linux Programming book, Michael Kerrisk</a>
    </li>
</ul>

<em>ET: 1h 33m</em>