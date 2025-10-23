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
pid:  1, ppid: 0, rss: 184 KB, name: init
pid:  1900, ppid: 479, rss: 21136 KB, name: go
pid:  1989, ppid: 1900, rss: 2392 KB, name: main
pid:  276, ppid: 1, rss: 212 KB, name: syslogd
pid:  304, ppid: 1, rss: 228 KB, name: crond
pid:  403, ppid: 1, rss: 168 KB, name: udhcpc
pid:  467, ppid: 1, rss: 180 KB, name: getty
pid:  479, ppid: 59, rss: 896 KB, name: sh
pid:  59, ppid: 1, rss: 10144 KB, name: orbstack-agent:
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

<em>ET: 30m</em>