var yzbPlayer = function(u) {
    var w = u.setOptions.width,
        x = u.setOptions.height,
        I = u.playerid,
        C = u.setOptions.poster,
        J = "",
        G = !0,
        K = u.setOptions.autostart ? u.setOptions.autostart : !1;
    if (void 0 == u.setOptions.width || "100%" == u.setOptions.width) w = document.documentElement.clientWidth;
    if (void 0 == u.setOptions.height || "100%" == u.setOptions.height) x = document.documentElement.clientHeight;
    void 0 != C && C.match(/.jpg|.jpeg|.png/i) || (C = "http://open.yizhibo.tv/player/poster.png");
    (function(c) {
        var a;
        a = window.XMLHttpRequest ? new XMLHttpRequest :
            new ActiveXObject("Microsoft.XMLHTTP");
        a.onreadystatechange = function() {
            if (4 == a.readyState)
                if (200 == a.status)
                    if (c.success) {
                        var f = a.getResponseHeader("Content-Type");
                        if (/application\/json/.test(f)) {
                            f = a.responseText;
                            try {
                                c.success(JSON.parse(f))
                            } catch (d) {
                                c.success({})
                            }
                        } else c.success(a.responseText)
                    } else console.log(a.responseText);
            else c.error ? c.error(a.status) : console.log("Error:" + a.status)
        };
        var l = "";
        if ("string" === typeof c.data) l = c.data,
            "?" == l.charAt(0) && (l = l.substring(1, l.length));
        else if ("object" ==
            typeof c.data) {
            for (key in c.data) l += key + "=" + c.data[key] + "&";
            l = l.replace(/&$/, "")
        }
        var f = c.type || "GET";
        "GET" == f ? (a.open(f, c.url + "?" + l, !1), a.send()) : (a.open(f, c.url, !0), a.setRequestHeader("Content-type", "application/x-www-form-urlencoded"), a.send(l))
    })({
        type: "GET",
        url: "http://open.yizhibo.tv/player/watchstart_relay.php",
        data: {
            app_nonce: u.app_nonce,
            vid: u.vid,
            name: u.name,
            app_key_id: u.app_key_id,
            app_signature: u.app_signature
        },
        success: function(c) {
            c = JSON.parse(c);
            var a = document.getElementById(I);
            "ok" == c.retval ?
                J = c.retinfo.play_url : (a.innerHTML = "E_LIVE_NOT_EXISTS" == c.retval ? '<div style="width:' + w + "px;background-color:#000;color:#fff;text-align:center;height:" + x + 'px;position:relative;z-index:9999;"><div class="loading" style="width:180px;height:225px;background:url(http://open.yizhibo.tv/player/load.gif) center center no-repeat;position:absolute;left:50%;top:50%;margin:-112px 0px 0px -90px;"></div></div>' : "E_API_INVALID_SIGNATURE" == c.retval ? '<div style="width:' + w + "px;background-color:#000;color:#fff;text-align:center;line-height:" +
                    x + 'px;position:relative;z-index:9999;">\u7b7e\u540d\u5931\u6548</div>' : "E_USER_NOT_EXISTS" == c.retval ? '<div style="width:' + w + "px;background-color:#000;color:#fff;text-align:center;line-height:" + x + 'px;position:relative;z-index:9999;">\u6613\u64ad\u53f7\u4e0d\u5b58\u5728</div>' : "E_API_INVALID_KEY_ID" == c.retval ? '<div style="width:' + w + "px;background-color:#000;color:#fff;text-align:center;line-height:" + x + 'px;position:relative;z-index:9999;">app_key_id\u9519\u8bef</div>' : '<div style="width:' + w + "px;background-color:#000;color:#fff;text-align:center;line-height:" +
                    x + 'px;position:relative;z-index:9999;">\u53c2\u6570\u9519\u8bef</div>', G = !1)
        },
        error: function() {
            L.innerHTML = '<div style="width:' + w + "px;background-color:#000;color:#fff;text-align:center;line-height:" + x + 'px;position:relative;z-index:9999;">\u8bf7\u6c42\u51fa\u73b0\u5f02\u5e38</div>';
            G = !1
        }
    });
    alert(u.app_nonce);
    if (G && 1 == G)
        if (u = navigator.userAgent.toLowerCase(), -1 < u.indexOf("android") || -1 < u.indexOf("Linux") || -1 < u.indexOf("iphone")) {
            var L = document.getElementById(I);
            u = document.createElement("video");
            u.src = J;
            u.width = w;
            u.height =
                x;
            u.setAttribute("webkit-playsinline", "webkit-playsinline");
            u.style.backgroundColor = "#000";
            1 == K && u.setAttribute("autoplay", K);
            u.setAttribute("controls", "controls");
            u.setAttribute("poster", C);
            L.appendChild(u, L.childNodes[0])
        } else "undefined" == typeof jwplayer && (jwplayer = function(c) {
                    if (jwplayer.api) return jwplayer.api.selectPlayer(c)
                },
                jwplayer.version = "6.8.4831", jwplayer.vid = document.createElement("video"), jwplayer.audio = document.createElement("audio"), jwplayer.source = document.createElement("source"),
                function(c) {
                    function a(p) {
                        return function() {
                            return q(p)
                        }
                    }

                    function l(p) {
                        return p && 0 <= p.indexOf("://") && p.split("/")[2] != g.location.href.split("/")[2]
                    }

                    function f(p, a, e) {
                        return function() {
                            p("Error loading file")
                        }
                    }

                    function h(p, a, e, b, f) {
                        return function() {
                            if (4 === p.readyState) switch (p.status) {
                                case 200:
                                    d(p, a, e, b, f)();
                                    break;
                                case 404:
                                    b("File not found")
                            }
                        }
                    }

                    function d(p, a, e, d, f) {
                        return function() {
                            var g,
                                m;
                            if (f) e(p);
                            else {
                                try {
                                    if (g = p.responseXML)
                                        if (m = g.firstChild, g.lastChild && "parsererror" === g.lastChild.nodeName) {
                                            d &&
                                                d("Invalid XML");
                                            return
                                        }
                                } catch (c) {}
                                if (g && m) return e(p);
                                (g = b.parseXML(p.responseText)) && g.firstChild ? (p = b.extend({},
                                    p, {
                                        responseXML: g
                                    }), e(p)) : d && d(p.responseText ? "Invalid XML" : a)
                            }
                        }
                    }
                    var n = document,
                        g = window,
                        m = navigator,
                        b = c.utils = function() {};
                    b.exists = function(p) {
                        switch (typeof p) {
                            case "string":
                                return 0 < p.length;
                            case "object":
                                return null !== p;
                            case "undefined":
                                return !1
                        }
                        return !0
                    };
                    b.styleDimension = function(p) {
                        return p + (0 < p.toString().indexOf("%") ? "" : "px")
                    };
                    b.getAbsolutePath = function(p, a) {
                        b.exists(a) || (a = n.location.href);
                        if (b.exists(p)) {
                            var e;
                            if (b.exists(p)) {
                                e = p.indexOf("://");
                                var g = p.indexOf("?");
                                e = 0 < e && (0 > g || g > e)
                            } else e = void 0;
                            if (e) return p;
                            e = a.substring(0, a.indexOf("://") + 3);
                            var g = a.substring(e.length, a.indexOf("/", e.length + 1)),
                                d;
                            0 === p.indexOf("/") ? d = p.split("/") : (d = a.split("?")[0], d = d.substring(e.length + g.length + 1, d.lastIndexOf("/")), d = d.split("/").concat(p.split("/")));
                            for (var f = [], m = 0; m < d.length; m++) d[m] && b.exists(d[m]) && "." != d[m] && (".." == d[m] ? f.pop() : f.push(d[m]));
                            return e + g + "/" + f.join("/")
                        }
                    };
                    b.extend = function() {
                        var a =
                            b.extend.arguments;
                        if (1 < a.length) {
                            for (var k = 1; k < a.length; k++) b.foreach(a[k],
                                function(e, k) {
                                    try {
                                        b.exists(k) && (a[0][e] = k)
                                    } catch (d) {}
                                });
                            return a[0]
                        }
                        return null
                    };
                    var r = window.console = window.console || {
                        log: function() {}
                    };
                    b.log = function() {
                        var a = Array.prototype.slice.call(arguments, 0);
                        "object" === typeof r.log ? r.log(a) : r.log.apply(r, a)
                    };
                    var q = b.userAgentMatch = function(a) {
                        return null !== m.userAgent.toLowerCase().match(a)
                    };
                    b.isIE = b.isMSIE = a(/msie/i);
                    b.isFF = a(/firefox/i);
                    b.isChrome = a(/chrome/i);
                    b.isIPod = a(/iP(hone|od)/i);
                    b.isIPad = a(/iPad/i);
                    b.isSafari602 = a(/Macintosh.*Mac OS X 10_8.*6\.0\.\d* Safari/i);
                    b.isIETrident = function(a) {
                        return a ? (a = parseFloat(a).toFixed(1), q(new RegExp("msie\\s*" + a + "|trident/.+rv:\\s*" + a, "i"))) : q(/msie|trident/i)
                    };
                    b.isSafari = function() {
                        return q(/safari/i) && !q(/chrome/i) && !q(/chromium/i) && !q(/android/i)
                    };
                    b.isIOS = function(a) {
                        return a ? q(new RegExp("iP(hone|ad|od).+\\sOS\\s" + a, "i")) : q(/iP(hone|ad|od)/i)
                    };
                    b.isAndroid = function(a, k) {
                        var e = k ? !q(/chrome\/[23456789]/i) : !0;
                        return a ? e && q(new RegExp("android.*" +
                            a, "i")) : e && q(/android/i)
                    };
                    b.isMobile = function() {
                        return b.isIOS() || b.isAndroid()
                    };
                    b.saveCookie = function(a, k) {
                        n.cookie = "jwplayer." + a + "=" + k + "; path=/"
                    };
                    b.getCookies = function() {
                        for (var a = {},
                                k = n.cookie.split("; "), e = 0; e < k.length; e++) {
                            var b = k[e].split("=");
                            0 === b[0].indexOf("jwplayer.") && (a[b[0].substring(9, b[0].length)] = b[1])
                        }
                        return a
                    };
                    b.typeOf = function(a) {
                        var b = typeof a;
                        return "object" === b ? a ? a instanceof Array ? "array" : b : "null" : b
                    };
                    b.translateEventResponse = function(a, k) {
                        var e = b.extend({},
                            k);
                        if (a != c.events.JWPLAYER_FULLSCREEN ||
                            e.fullscreen)
                            if ("object" == typeof e.data) {
                                var d = e.data;
                                delete e.data;
                                e = b.extend(e, d)
                            } else "object" == typeof e.metadata && b.deepReplaceKeyName(e.metadata, ["__dot__", "__spc__", "__dsh__", "__default__"], [".", " ", "-", "default"]);
                        else e.fullscreen = "true" == e.message ? !0 : !1,
                            delete e.message;
                        b.foreach(["position", "duration", "offset"],
                            function(a, b) {
                                e[b] && (e[b] = Math.round(1E3 * e[b]) / 1E3)
                            });
                        return e
                    };
                    b.flashVersion = function() {
                        if (b.isAndroid()) return 0;
                        var a = m.plugins,
                            k;
                        try {
                            if ("undefined" !== a && (k = a["Shockwave Flash"])) return parseInt(k.description.replace(/\D+(\d+)\..*/,
                                "$1"), 10)
                        } catch (e) {}
                        if ("undefined" != typeof g.ActiveXObject) try {
                            if (k = new g.ActiveXObject("ShockwaveFlash.ShockwaveFlash")) return parseInt(k.GetVariable("$version").split(" ")[1].split(",")[0], 10)
                        } catch (e) {}
                        return 0
                    };
                    b.getScriptPath = function(a) {
                        for (var b = n.getElementsByTagName("script"), e = 0; e < b.length; e++) {
                            var d = b[e].src;
                            if (d && 0 <= d.indexOf(a)) return d.substr(0, d.indexOf(a))
                        }
                        return ""
                    };
                    b.deepReplaceKeyName = function(a, k, e) {
                        switch (c.utils.typeOf(a)) {
                            case "array":
                                for (var d = 0; d < a.length; d++) a[d] = c.utils.deepReplaceKeyName(a[d],
                                    k, e);
                                break;
                            case "object":
                                b.foreach(a,
                                    function(b, d) {
                                        var g;
                                        if (k instanceof Array && e instanceof Array) {
                                            if (k.length != e.length) return;
                                            g = k
                                        } else g = [k];
                                        for (var f = b, m = 0; m < g.length; m++) f = f.replace(new RegExp(k[m], "g"), e[m]);
                                        a[f] = c.utils.deepReplaceKeyName(d, k, e);
                                        b != f && delete a[b]
                                    })
                        }
                        return a
                    };
                    var t = b.pluginPathType = {
                        ABSOLUTE: 0,
                        RELATIVE: 1,
                        CDN: 2
                    };
                    b.getPluginPathType = function(a) {
                        if ("string" == typeof a) {
                            a = a.split("?")[0];
                            var k = a.indexOf("://");
                            if (0 < k) return t.ABSOLUTE;
                            var e = a.indexOf("/");
                            a = b.extension(a);
                            return !(0 >
                                k && 0 > e) || a && isNaN(a) ? t.RELATIVE : t.CDN
                        }
                    };
                    b.getPluginName = function(a) {
                        return a.replace(/^(.*\/)?([^-]*)-?.*\.(swf|js)$/, "$2")
                    };
                    b.getPluginVersion = function(a) {
                        return a.replace(/[^-]*-?([^\.]*).*$/, "$1")
                    };
                    b.isYouTube = function(a) {
                        return /^(http|\/\/).*(youtube\.com|youtu\.be)\/.+/.test(a)
                    };
                    b.youTubeID = function(a) {
                        try {
                            return /v[=\/]([^?&]*)|youtu\.be\/([^?]*)|^([\w-]*)$/i.exec(a).slice(1).join("").replace("?", "")
                        } catch (b) {
                            return ""
                        }
                    };
                    b.isRtmp = function(a, b) {
                        return 0 === a.indexOf("rtmp") || "rtmp" == b
                    };
                    b.foreach =
                        function(a, k) {
                            var e,
                                d;
                            for (e in a) "function" == b.typeOf(a.hasOwnProperty) ? a.hasOwnProperty(e) && (d = a[e], k(e, d)) : (d = a[e], k(e, d))
                        };
                    b.isHTTPS = function() {
                        return 0 === g.location.href.indexOf("https")
                    };
                    b.repo = function() {
                        var a = "";
                        try {
                            b.isHTTPS() && (a = a.replace("http://", "https://ssl."))
                        } catch (k) {}
                        return a
                    };
                    b.ajax = function(a, k, e, m) {
                        var c;
                        0 < a.indexOf("#") && (a = a.replace(/#.*$/, ""));
                        if (l(a) && b.exists(g.XDomainRequest)) c = new g.XDomainRequest,
                            c.onload = d(c, a, k, e, m),
                            c.ontimeout = c.onprogress = function() {},
                            c.timeout = 5E3;
                        else if (b.exists(g.XMLHttpRequest)) c = new g.XMLHttpRequest,
                            c.onreadystatechange = h(c, a, k, e, m);
                        else return e && e(),
                            c;
                        c.overrideMimeType && c.overrideMimeType("text/xml");
                        c.onerror = f(e, a, c);
                        setTimeout(function() {
                                try {
                                    c.open("GET", a, !0),
                                        c.send()
                                } catch (b) {
                                    e && e(a)
                                }
                            },
                            0);
                        return c
                    };
                    b.parseXML = function(a) {
                        var b;
                        try {
                            if (g.DOMParser) {
                                if (b = (new g.DOMParser).parseFromString(a, "text/xml"), b.childNodes && b.childNodes.length && "parsererror" == b.childNodes[0].firstChild.nodeName) return
                            } else b = new g.ActiveXObject("Microsoft.XMLDOM"),
                                b.async = "false",
                                b.loadXML(a)
                        } catch (e) {
                            return
                        }
                        return b
                    };
                    b.filterPlaylist = function(a, k) {
                        var e = [],
                            d,
                            g,
                            m,
                            f;
                        for (d = 0; d < a.length; d++)
                            if (g = b.extend({},
                                    a[d]), g.sources = b.filterSources(g.sources), 0 < g.sources.length) {
                                for (m = 0; m < g.sources.length; m++) f = g.sources[m],
                                    f.label || (f.label = m.toString());
                                e.push(g)
                            }
                        if (k && 0 === e.length)
                            for (d = 0; d < a.length; d++)
                                if (g = b.extend({},
                                        a[d]), g.sources = b.filterSources(g.sources, !0), 0 < g.sources.length) {
                                    for (m = 0; m < g.sources.length; m++) f = g.sources[m],
                                        f.label || (f.label = m.toString());
                                    e.push(g)
                                }
                        return e
                    };
                    b.filterSources = function(a, d) {
                        var e,
                            g,
                            m = b.extensionmap;
                        if (a) {
                            g = [];
                            for (var f = 0; f < a.length; f++) {
                                var n = a[f].type,
                                    h = a[f].file;
                                h && (h = b.trim(h));
                                n || (n = m.extType(b.extension(h)), a[f].type = n);
                                d ? c.embed.flashCanPlay(h, n) && (e || (e = n), n == e && g.push(b.extend({},
                                    a[f]))) : b.canPlayHTML5(n) && (e || (e = n), n == e && g.push(b.extend({},
                                    a[f])))
                            }
                        }
                        return g
                    };
                    b.canPlayHTML5 = function(a) {
                        if (b.isAndroid() && ("hls" == a || "m3u" == a || "m3u8" == a)) return !1;
                        a = b.extensionmap.types[a];
                        return !!a && !!c.vid.canPlayType && c.vid.canPlayType(a)
                    };
                    b.seconds =
                        function(a) {
                            a = a.replace(",", ".");
                            var b = a.split(":"),
                                e = 0;
                            "s" == a.slice(-1) ? e = parseFloat(a) : "m" == a.slice(-1) ? e = 60 * parseFloat(a) : "h" == a.slice(-1) ? e = 3600 * parseFloat(a) : 1 < b.length ? (e = parseFloat(b[b.length - 1]), e += 60 * parseFloat(b[b.length - 2]), 3 == b.length && (e += 3600 * parseFloat(b[b.length - 3]))) : e = parseFloat(a);
                            return e
                        };
                    b.serialize = function(a) {
                        return null == a ? null : "true" == a.toString().toLowerCase() ? !0 : "false" == a.toString().toLowerCase() ? !1 : isNaN(Number(a)) || 5 < a.length || 0 === a.length ? a : Number(a)
                    }
                }(jwplayer),
                function(c) {
                    var a =
                        "video/",
                        l = c.foreach,
                        f = {
                            mp4: a + "mp4",
                            vorbis: "audio/ogg",
                            ogg: a + "ogg",
                            webm: a + "webm",
                            aac: "audio/mp4",
                            mp3: "audio/mpeg",
                            hls: "application/vnd.apple.mpegurl"
                        },
                        h = {
                            mp4: f.mp4,
                            f4v: f.mp4,
                            m4v: f.mp4,
                            mov: f.mp4,
                            m4a: f.aac,
                            f4a: f.aac,
                            aac: f.aac,
                            mp3: f.mp3,
                            ogv: f.ogg,
                            ogg: f.vorbis,
                            oga: f.vorbis,
                            webm: f.webm,
                            m3u8: f.hls,
                            hls: f.hls
                        },
                        a = "video",
                        a = {
                            flv: a,
                            f4v: a,
                            mov: a,
                            m4a: a,
                            m4v: a,
                            mp4: a,
                            aac: a,
                            f4a: a,
                            mp3: "sound",
                            smil: "rtmp",
                            m3u8: "hls",
                            hls: "hls"
                        },
                        d = c.extensionmap = {};
                    l(h,
                        function(a, g) {
                            d[a] = {
                                html5: g
                            }
                        });
                    l(a,
                        function(a, g) {
                            d[a] || (d[a] = {});
                            d[a].flash =
                                g
                        });
                    d.types = f;
                    d.mimeType = function(a) {
                        var d;
                        l(f,
                            function(f, b) {
                                d || b != a || (d = f)
                            });
                        return d
                    };
                    d.extType = function(a) {
                        return d.mimeType(h[a])
                    }
                }(jwplayer.utils),
                function(c) {
                    var a = c.loaderstatus = {
                            NEW: 0,
                            LOADING: 1,
                            ERROR: 2,
                            COMPLETE: 3
                        },
                        l = document;
                    c.scriptloader = function(f) {
                        function h(b) {
                            n = a.ERROR;
                            m.sendEvent(g.ERROR)
                        }

                        function d(b) {
                            n = a.COMPLETE;
                            m.sendEvent(g.COMPLETE)
                        }
                        var n = a.NEW,
                            g = jwplayer.events,
                            m = new g.eventdispatcher;
                        c.extend(this, m);
                        this.load = function() {
                            var b = c.scriptloader.loaders[f];
                            if (!b || b.getStatus() !=
                                a.NEW && b.getStatus() != a.LOADING) {
                                if (c.scriptloader.loaders[f] = this, n == a.NEW) {
                                    n = a.LOADING;
                                    var m = l.createElement("script");
                                    m.addEventListener ? (m.onload = d, m.onerror = h) : m.readyState && (m.onreadystatechange = function() {
                                        "loaded" != m.readyState && "complete" != m.readyState || d()
                                    });
                                    l.getElementsByTagName("head")[0].appendChild(m);
                                    m.src = f
                                }
                            } else b.addEventListener(g.ERROR, h),
                                b.addEventListener(g.COMPLETE, d)
                        };
                        this.getStatus = function() {
                            return n
                        }
                    };
                    c.scriptloader.loaders = {}
                }(jwplayer.utils),
                function(c) {
                    c.trim = function(a) {
                        return a.replace(/^\s*/,
                            "").replace(/\s*$/, "")
                    };
                    c.pad = function(a, c, f) {
                        for (f || (f = "0"); a.length < c;) a = f + a;
                        return a
                    };
                    c.xmlAttribute = function(a, c) {
                        for (var f = 0; f < a.attributes.length; f++)
                            if (a.attributes[f].name && a.attributes[f].name.toLowerCase() == c.toLowerCase()) return a.attributes[f].value.toString();
                        return ""
                    };
                    c.extension = function(a) {
                        if (!a || "rtmp" == a.substr(0, 4)) return "";
                        a = a.substring(a.lastIndexOf("/") + 1, a.length).split("?")[0].split("#")[0];
                        if (-1 < a.lastIndexOf(".")) return a.substr(a.lastIndexOf(".") + 1, a.length).toLowerCase()
                    };
                    c.stringToColor = function(a) {
                        a = a.replace(/(#|0x)?([0-9A-F]{3,6})$/gi, "$2");
                        3 == a.length && (a = a.charAt(0) + a.charAt(0) + a.charAt(1) + a.charAt(1) + a.charAt(2) + a.charAt(2));
                        return parseInt(a, 16)
                    }
                }(jwplayer.utils),
                function(c) {
                    c.touch = function(a) {
                        function l(a) {
                            "touchstart" == a.type ? (n = !0, m = h(r.DRAG_START, a)) : "touchmove" == a.type ? n && (b || (f(r.DRAG_START, a, m), b = !0), f(r.DRAG, a)) : (n && (b ? f(r.DRAG_END, a) : (a.cancelBubble = !0, f(r.TAP, a))), n = b = !1, m = null)
                        }

                        function f(a, b, d) {
                            if (g[a] && (b.preventManipulation && b.preventManipulation(),
                                    b.preventDefault && b.preventDefault(), b = d ? d : h(a, b))) g[a](b)
                        }

                        function h(a, b) {
                            var g = null;
                            b.touches && b.touches.length ? g = b.touches[0] : b.changedTouches && b.changedTouches.length && (g = b.changedTouches[0]);
                            if (!g) return null;
                            var k = d.getBoundingClientRect(),
                                g = {
                                    type: a,
                                    target: d,
                                    x: g.pageX - window.pageXOffset - k.left,
                                    y: g.pageY,
                                    deltaX: 0,
                                    deltaY: 0
                                };
                            a != r.TAP && m && (g.deltaX = g.x - m.x, g.deltaY = g.y - m.y);
                            return g
                        }
                        var d = a,
                            n = !1,
                            g = {},
                            m = null,
                            b = !1,
                            r = c.touchEvents;
                        document.addEventListener("touchmove", l);
                        document.addEventListener("touchend",
                            function(a) {
                                n && b && f(r.DRAG_END, a);
                                n = b = !1;
                                m = null
                            });
                        document.addEventListener("touchcancel", l);
                        a.addEventListener("touchstart", l);
                        a.addEventListener("touchend", l);
                        this.addEventListener = function(a, b) {
                            g[a] = b
                        };
                        this.removeEventListener = function(a) {
                            delete g[a]
                        };
                        return this
                    }
                }(jwplayer.utils), jwplayer.utils.touchEvents = {
                    DRAG: "jwplayerDrag",
                    DRAG_START: "jwplayerDragStart",
                    DRAG_END: "jwplayerDragEnd",
                    TAP: "jwplayerTap"
                },
                jwplayer.events = {
                    COMPLETE: "COMPLETE",
                    ERROR: "ERROR",
                    API_READY: "jwplayerAPIReady",
                    JWPLAYER_READY: "jwplayerReady",
                    JWPLAYER_FULLSCREEN: "jwplayerFullscreen",
                    JWPLAYER_RESIZE: "jwplayerResize",
                    JWPLAYER_ERROR: "jwplayerError",
                    JWPLAYER_SETUP_ERROR: "jwplayerSetupError",
                    JWPLAYER_MEDIA_BEFOREPLAY: "jwplayerMediaBeforePlay",
                    JWPLAYER_MEDIA_BEFORECOMPLETE: "jwplayerMediaBeforeComplete",
                    JWPLAYER_COMPONENT_SHOW: "jwplayerComponentShow",
                    JWPLAYER_COMPONENT_HIDE: "jwplayerComponentHide",
                    JWPLAYER_MEDIA_BUFFER: "jwplayerMediaBuffer",
                    JWPLAYER_MEDIA_BUFFER_FULL: "jwplayerMediaBufferFull",
                    JWPLAYER_MEDIA_ERROR: "jwplayerMediaError",
                    JWPLAYER_MEDIA_LOADED: "jwplayerMediaLoaded",
                    JWPLAYER_MEDIA_COMPLETE: "jwplayerMediaComplete",
                    JWPLAYER_MEDIA_SEEK: "jwplayerMediaSeek",
                    JWPLAYER_MEDIA_TIME: "jwplayerMediaTime",
                    JWPLAYER_MEDIA_VOLUME: "jwplayerMediaVolume",
                    JWPLAYER_MEDIA_META: "jwplayerMediaMeta",
                    JWPLAYER_MEDIA_MUTE: "jwplayerMediaMute",
                    JWPLAYER_MEDIA_LEVELS: "jwplayerMediaLevels",
                    JWPLAYER_MEDIA_LEVEL_CHANGED: "jwplayerMediaLevelChanged",
                    JWPLAYER_CAPTIONS_CHANGED: "jwplayerCaptionsChanged",
                    JWPLAYER_CAPTIONS_LIST: "jwplayerCaptionsList",
                    JWPLAYER_CAPTIONS_LOADED: "jwplayerCaptionsLoaded",
                    JWPLAYER_PLAYER_STATE: "jwplayerPlayerState",
                    state: {
                        BUFFERING: "BUFFERING",
                        IDLE: "IDLE",
                        PAUSED: "PAUSED",
                        PLAYING: "PLAYING"
                    },
                    JWPLAYER_PLAYLIST_LOADED: "jwplayerPlaylistLoaded",
                    JWPLAYER_PLAYLIST_ITEM: "jwplayerPlaylistItem",
                    JWPLAYER_PLAYLIST_COMPLETE: "jwplayerPlaylistComplete",
                    JWPLAYER_DISPLAY_CLICK: "jwplayerViewClick",
                    JWPLAYER_CONTROLS: "jwplayerViewControls",
                    JWPLAYER_USER_ACTION: "jwplayerUserAction",
                    JWPLAYER_INSTREAM_CLICK: "jwplayerInstreamClicked",
                    JWPLAYER_INSTREAM_DESTROYED: "jwplayerInstreamDestroyed",
                    JWPLAYER_AD_TIME: "jwplayerAdTime",
                    JWPLAYER_AD_ERROR: "jwplayerAdError",
                    JWPLAYER_AD_CLICK: "jwplayerAdClicked",
                    JWPLAYER_AD_COMPLETE: "jwplayerAdComplete",
                    JWPLAYER_AD_IMPRESSION: "jwplayerAdImpression",
                    JWPLAYER_AD_COMPANIONS: "jwplayerAdCompanions",
                    JWPLAYER_AD_SKIPPED: "jwplayerAdSkipped"
                },
                function(c) {
                    var a = c.utils;
                    c.events.eventdispatcher = function(l, f) {
                        function h(d, m, b) {
                            if (d)
                                for (var f = 0; f < d.length; f++) {
                                    var c = d[f];
                                    if (c) {
                                        null !== c.count && 0 === --c.count && delete d[f];
                                        try {
                                            c.listener(m)
                                        } catch (n) {
                                            a.log('Error handling "' +
                                                b + '" event listener [' + f + "]: " + n.toString(), c.listener, m)
                                        }
                                    }
                                }
                        }
                        var d,
                            n;
                        this.resetEventListeners = function() {
                            d = {};
                            n = []
                        };
                        this.resetEventListeners();
                        this.addEventListener = function(g, f, b) {
                            try {
                                a.exists(d[g]) || (d[g] = []),
                                    "string" == a.typeOf(f) && (f = (new Function("return " + f))()),
                                    d[g].push({
                                        listener: f,
                                        count: b || null
                                    })
                            } catch (c) {
                                a.log("error", c)
                            }
                            return !1
                        };
                        this.removeEventListener = function(f, m) {
                            if (d[f]) {
                                try {
                                    for (var b = 0; b < d[f].length; b++)
                                        if (d[f][b].listener.toString() == m.toString()) {
                                            d[f].splice(b, 1);
                                            break
                                        }
                                } catch (c) {
                                    a.log("error",
                                        c)
                                }
                                return !1
                            }
                        };
                        this.addGlobalListener = function(d, f) {
                            try {
                                "string" == a.typeOf(d) && (d = (new Function("return " + d))()),
                                    n.push({
                                        listener: d,
                                        count: f || null
                                    })
                            } catch (b) {
                                a.log("error", b)
                            }
                            return !1
                        };
                        this.removeGlobalListener = function(d) {
                            if (d) {
                                try {
                                    for (var f = n.length; f--;) n[f].listener.toString() == d.toString() && n.splice(f, 1)
                                } catch (b) {
                                    a.log("error", b)
                                }
                                return !1
                            }
                        };
                        this.sendEvent = function(g, m) {
                            a.exists(m) || (m = {});
                            a.extend(m, {
                                id: l,
                                version: c.version,
                                type: g
                            });
                            f && a.log(g, m);
                            h(d[g], m, g);
                            h(n, m, g)
                        }
                    }
                }(window.jwplayer),
                function(c) {
                    var a = {},
                        l = {};
                    c.plugins = function() {};
                    c.plugins.loadPlugins = function(f, h) {
                        l[f] = new c.plugins.pluginloader(new c.plugins.model(a), h);
                        return l[f]
                    };
                    c.plugins.registerPlugin = function(f, h, d, n) {
                        var g = c.utils.getPluginName(f);
                        a[g] || (a[g] = new c.plugins.plugin(f));
                        a[g].registerPlugin(f, h, d, n)
                    }
                }(jwplayer),
                function(c) {
                    c.plugins.model = function(a) {
                        this.addPlugin = function(l) {
                            var f = c.utils.getPluginName(l);
                            a[f] || (a[f] = new c.plugins.plugin(l));
                            return a[f]
                        };
                        this.getPlugins = function() {
                            return a
                        }
                    }
                }(jwplayer),
                function(c) {
                    var a =
                        jwplayer.utils,
                        l = jwplayer.events;
                    c.pluginmodes = {
                        FLASH: 0,
                        JAVASCRIPT: 1,
                        HYBRID: 2
                    };
                    c.plugin = function(f) {
                        function h() {
                            switch (a.getPluginPathType(f)) {
                                case a.pluginPathType.ABSOLUTE:
                                    return f;
                                case a.pluginPathType.RELATIVE:
                                    return a.getAbsolutePath(f, window.location.href)
                            }
                        }

                        function d(b) {
                            q = setTimeout(function() {
                                    g = a.loaderstatus.COMPLETE;
                                    t.sendEvent(l.COMPLETE)
                                },
                                1E3)
                        }

                        function n(b) {
                            g = a.loaderstatus.ERROR;
                            t.sendEvent(l.ERROR)
                        }
                        var g = a.loaderstatus.NEW,
                            m,
                            b,
                            r,
                            q,
                            t = new l.eventdispatcher;
                        a.extend(this, t);
                        this.load =
                            function() {
                                if (g == a.loaderstatus.NEW)
                                    if (0 < f.lastIndexOf(".swf")) m = f,
                                        g = a.loaderstatus.COMPLETE,
                                        t.sendEvent(l.COMPLETE);
                                    else if (a.getPluginPathType(f) == a.pluginPathType.CDN) g = a.loaderstatus.COMPLETE,
                                    t.sendEvent(l.COMPLETE);
                                else {
                                    g = a.loaderstatus.LOADING;
                                    var b = new a.scriptloader(h());
                                    b.addEventListener(l.COMPLETE, d);
                                    b.addEventListener(l.ERROR, n);
                                    b.load()
                                }
                            };
                        this.registerPlugin = function(d, f, e, c) {
                            q && (clearTimeout(q), q = void 0);
                            r = f;
                            e && c ? (m = c, b = e) : "string" == typeof e ? m = e : "function" == typeof e ? b = e : e || c || (m = d);
                            g = a.loaderstatus.COMPLETE;
                            t.sendEvent(l.COMPLETE)
                        };
                        this.getStatus = function() {
                            return g
                        };
                        this.getPluginName = function() {
                            return a.getPluginName(f)
                        };
                        this.getFlashPath = function() {
                            if (m) switch (a.getPluginPathType(m)) {
                                case a.pluginPathType.ABSOLUTE:
                                    return m;
                                case a.pluginPathType.RELATIVE:
                                    return 0 < f.lastIndexOf(".swf") ? a.getAbsolutePath(m, window.location.href) : a.getAbsolutePath(m, h())
                            }
                            return null
                        };
                        this.getJS = function() {
                            return b
                        };
                        this.getTarget = function() {
                            return r
                        };
                        this.getPluginmode = function() {
                            if ("undefined" !=
                                typeof m && "undefined" != typeof b) return c.pluginmodes.HYBRID;
                            if ("undefined" != typeof m) return c.pluginmodes.FLASH;
                            if ("undefined" != typeof b) return c.pluginmodes.JAVASCRIPT
                        };
                        this.getNewInstance = function(a, d, e) {
                            return new b(a, d, e)
                        };
                        this.getURL = function() {
                            return f
                        }
                    }
                }(jwplayer.plugins),
                function(c) {
                    var a = c.utils,
                        l = c.events,
                        f = a.foreach;
                    c.plugins.pluginloader = function(h, d) {
                        function n() {
                            r ? p.sendEvent(l.ERROR, {
                                message: q
                            }) : b || (b = !0, m = a.loaderstatus.COMPLETE, p.sendEvent(l.COMPLETE))
                        }

                        function g() {
                            t || n();
                            if (!b &&
                                !r) {
                                var e = 0,
                                    d = h.getPlugins();
                                a.foreach(t,
                                    function(b, f) {
                                        var k = a.getPluginName(b),
                                            g = d[k],
                                            k = g.getJS(),
                                            m = g.getTarget(),
                                            g = g.getStatus();
                                        g == a.loaderstatus.LOADING || g == a.loaderstatus.NEW ? e++ : k && (!m || parseFloat(m) > parseFloat(c.version)) && (r = !0, q = "Incompatible player version", n())
                                    });
                                0 == e && n()
                            }
                        }
                        var m = a.loaderstatus.NEW,
                            b = !1,
                            r = !1,
                            q,
                            t = d,
                            p = new l.eventdispatcher;
                        a.extend(this, p);
                        this.setupPlugins = function(e, b, d) {
                            var k = {
                                    length: 0,
                                    plugins: {}
                                },
                                g = 0,
                                c = {},
                                m = h.getPlugins();
                            f(b.plugins,
                                function(f, n) {
                                    var h = a.getPluginName(f),
                                        l = m[h],
                                        F = l.getFlashPath(),
                                        p = l.getJS(),
                                        r = l.getURL();
                                    F && (k.plugins[F] = a.extend({},
                                        n), k.plugins[F].pluginmode = l.getPluginmode(), k.length++);
                                    try {
                                        if (p && b.plugins && b.plugins[r]) {
                                            var q = document.createElement("div");
                                            q.id = e.id + "_" + h;
                                            q.style.position = "absolute";
                                            q.style.top = 0;
                                            q.style.zIndex = g + 10;
                                            c[h] = l.getNewInstance(e, a.extend({},
                                                b.plugins[r]), q);
                                            g++;
                                            e.onReady(d(c[h], q, !0));
                                            e.onResize(d(c[h], q))
                                        }
                                    } catch (t) {
                                        a.log("ERROR: Failed to load " + h + ".")
                                    }
                                });
                            e.plugins = c;
                            return k
                        };
                        this.load = function() {
                            if (!a.exists(d) || "object" ==
                                a.typeOf(d)) {
                                m = a.loaderstatus.LOADING;
                                f(d,
                                    function(b, e) {
                                        if (a.exists(b)) {
                                            var d = h.addPlugin(b);
                                            d.addEventListener(l.COMPLETE, g);
                                            d.addEventListener(l.ERROR, k)
                                        }
                                    });
                                var b = h.getPlugins();
                                f(b,
                                    function(a, b) {
                                        b.load()
                                    })
                            }
                            g()
                        };
                        var k = this.pluginFailed = function(a) {
                            r || (r = !0, q = "File not found", n())
                        };
                        this.getStatus = function() {
                            return m
                        }
                    }
                }(jwplayer),
                function(c) {
                    jwplayer.parsers = {
                        localName: function(a) {
                            return a ? a.localName ? a.localName : a.baseName ? a.baseName : "" : ""
                        },
                        textContent: function(a) {
                            return a ? a.textContent ? jwplayer.utils.trim(a.textContent) :
                                a.text ? jwplayer.utils.trim(a.text) : "" : ""
                        },
                        getChildNode: function(a, c) {
                            return a.childNodes[c]
                        },
                        numChildren: function(a) {
                            return a.childNodes ? a.childNodes.length : 0
                        }
                    }
                }(jwplayer),
                function(c) {
                    var a = c.parsers;
                    (a.jwparser = function() {}).parseEntry = function(l, f) {
                        for (var h = [], d = [], n = c.utils.xmlAttribute, g = 0; g < l.childNodes.length; g++) {
                            var m = l.childNodes[g];
                            if ("jwplayer" == m.prefix) {
                                var b = a.localName(m);
                                "source" == b ? (delete f.sources, h.push({
                                        file: n(m, "file"),
                                        "default": n(m, "default"),
                                        label: n(m, "label"),
                                        type: n(m, "type")
                                    })) :
                                    "track" == b ? (delete f.tracks, d.push({
                                        file: n(m, "file"),
                                        "default": n(m, "default"),
                                        kind: n(m, "kind"),
                                        label: n(m, "label")
                                    })) : (f[b] = c.utils.serialize(a.textContent(m)), "file" == b && f.sources && delete f.sources)
                            }
                            f.file || (f.file = f.link)
                        }
                        if (h.length)
                            for (f.sources = [], g = 0; g < h.length; g++) 0 < h[g].file.length && (h[g]["default"] = "true" == h[g]["default"] ? !0 : !1, h[g].label.length || delete h[g].label, f.sources.push(h[g]));
                        if (d.length)
                            for (f.tracks = [], g = 0; g < d.length; g++) 0 < d[g].file.length && (d[g]["default"] = "true" == d[g]["default"] ?
                                !0 : !1, d[g].kind = d[g].kind.length ? d[g].kind : "captions", d[g].label.length || delete d[g].label, f.tracks.push(d[g]));
                        return f
                    }
                }(jwplayer),
                function(c) {
                    var a = jwplayer.utils,
                        l = a.xmlAttribute,
                        f = c.localName,
                        h = c.textContent,
                        d = c.numChildren,
                        n = c.mediaparser = function() {};
                    n.parseGroup = function(g, c) {
                        var b,
                            r,
                            q = [];
                        for (r = 0; r < d(g); r++)
                            if (b = g.childNodes[r], "media" == b.prefix && f(b)) switch (f(b).toLowerCase()) {
                                case "content":
                                    l(b, "duration") && (c.duration = a.seconds(l(b, "duration")));
                                    0 < d(b) && (c = n.parseGroup(b, c));
                                    l(b, "url") &&
                                        (c.sources || (c.sources = []), c.sources.push({
                                            file: l(b, "url"),
                                            type: l(b, "type"),
                                            width: l(b, "width"),
                                            label: l(b, "label")
                                        }));
                                    break;
                                case "title":
                                    c.title = h(b);
                                    break;
                                case "description":
                                    c.description = h(b);
                                    break;
                                case "guid":
                                    c.mediaid = h(b);
                                    break;
                                case "thumbnail":
                                    c.image || (c.image = l(b, "url"));
                                    break;
                                case "group":
                                    n.parseGroup(b, c);
                                    break;
                                case "subtitle":
                                    var t = {};
                                    t.file = l(b, "url");
                                    t.kind = "captions";
                                    if (0 < l(b, "lang").length) {
                                        var p = t;
                                        b = l(b, "lang");
                                        var k = {
                                            zh: "Chinese",
                                            nl: "Dutch",
                                            en: "English",
                                            fr: "French",
                                            de: "German",
                                            it: "Italian",
                                            ja: "Japanese",
                                            pt: "Portuguese",
                                            ru: "Russian",
                                            es: "Spanish"
                                        };
                                        b = k[b] ? k[b] : b;
                                        p.label = b
                                    }
                                    q.push(t)
                            }
                        c.hasOwnProperty("tracks") || (c.tracks = []);
                        for (r = 0; r < q.length; r++) c.tracks.push(q[r]);
                        return c
                    }
                }(jwplayer.parsers),
                function(c) {
                    function a(a) {
                        for (var d = {},
                                b = 0; b < a.childNodes.length; b++) {
                            var h = a.childNodes[b],
                                q = n(h);
                            if (q) switch (q.toLowerCase()) {
                                case "enclosure":
                                    d.file = l.xmlAttribute(h, "url");
                                    break;
                                case "title":
                                    d.title = f(h);
                                    break;
                                case "guid":
                                    d.mediaid = f(h);
                                    break;
                                case "pubdate":
                                    d.date = f(h);
                                    break;
                                case "description":
                                    d.description =
                                        f(h);
                                    break;
                                case "link":
                                    d.link = f(h);
                                    break;
                                case "category":
                                    d.tags = d.tags ? d.tags + f(h) : f(h)
                            }
                        }
                        d = c.mediaparser.parseGroup(a, d);
                        d = c.jwparser.parseEntry(a, d);
                        return new jwplayer.playlist.item(d)
                    }
                    var l = jwplayer.utils,
                        f = c.textContent,
                        h = c.getChildNode,
                        d = c.numChildren,
                        n = c.localName;
                    c.rssparser = {};
                    c.rssparser.parse = function(f) {
                        for (var c = [], b = 0; b < d(f); b++) {
                            var l = h(f, b);
                            if ("channel" == n(l).toLowerCase())
                                for (var q = 0; q < d(l); q++) {
                                    var t = h(l, q);
                                    "item" == n(t).toLowerCase() && c.push(a(t))
                                }
                        }
                        return c
                    }
                }(jwplayer.parsers),
                function(c) {
                    c.playlist =
                        function(a) {
                            var l = [];
                            if ("array" == c.utils.typeOf(a))
                                for (var f = 0; f < a.length; f++) l.push(new c.playlist.item(a[f]));
                            else l.push(new c.playlist.item(a));
                            return l
                        }
                }(jwplayer),
                function(c) {
                    var a = c.item = function(l) {
                        var f = jwplayer.utils,
                            h = f.extend({},
                                a.defaults, l);
                        h.tracks = l && f.exists(l.tracks) ? l.tracks : [];
                        0 == h.sources.length && (h.sources = [new c.source(h)]);
                        for (var d = 0; d < h.sources.length; d++) {
                            var n = h.sources[d]["default"];
                            h.sources[d]["default"] = n ? "true" == n.toString() : !1;
                            h.sources[d] = new c.source(h.sources[d])
                        }
                        if (h.captions &&
                            !f.exists(l.tracks)) {
                            for (l = 0; l < h.captions.length; l++) h.tracks.push(h.captions[l]);
                            delete h.captions
                        }
                        for (d = 0; d < h.tracks.length; d++) h.tracks[d] = new c.track(h.tracks[d]);
                        return h
                    };
                    a.defaults = {
                        description: "",
                        image: "",
                        mediaid: "",
                        title: "",
                        sources: [],
                        tracks: []
                    }
                }(jwplayer.playlist),
                function(c) {
                    var a = jwplayer,
                        l = a.utils,
                        f = a.events,
                        h = a.parsers;
                    c.loader = function() {
                        function a(b) {
                            try {
                                var d = b.responseXML.childNodes;
                                b = "";
                                for (var n = 0; n < d.length && (b = d[n], 8 == b.nodeType); n++);
                                "xml" == h.localName(b) && (b = b.nextSibling);
                                if ("rss" != h.localName(b)) g("Not a valid RSS feed");
                                else {
                                    var l = new c(h.rssparser.parse(b));
                                    m.sendEvent(f.JWPLAYER_PLAYLIST_LOADED, {
                                        playlist: l
                                    })
                                }
                            } catch (p) {
                                g()
                            }
                        }

                        function n(a) {
                            g(a.match(/invalid/i) ? "Not a valid RSS feed" : "")
                        }

                        function g(a) {
                            m.sendEvent(f.JWPLAYER_ERROR, {
                                message: a ? a : "\u89c6\u9891\u6587\u4ef6\u52a0\u8f7d\u9519\u8bef"
                            })
                        }
                        var m = new f.eventdispatcher;
                        l.extend(this, m);
                        this.load = function(b) {
                            l.ajax(b, a, n)
                        }
                    }
                }(jwplayer.playlist),
                function(c) {
                    var a = jwplayer.utils,
                        l = {
                            file: void 0,
                            label: void 0,
                            type: void 0,
                            "default": void 0
                        };
                    c.source = function(f) {
                        var c = a.extend({},
                            l);
                        a.foreach(l,
                            function(d, n) {
                                a.exists(f[d]) && (c[d] = f[d], delete f[d])
                            });
                        c.type && 0 < c.type.indexOf("/") && (c.type = a.extensionmap.mimeType(c.type));
                        "m3u8" == c.type && (c.type = "hls");
                        "smil" == c.type && (c.type = "rtmp");
                        return c
                    }
                }(jwplayer.playlist),
                function(c) {
                    var a = jwplayer.utils,
                        l = {
                            file: void 0,
                            label: void 0,
                            kind: "captions",
                            "default": !1
                        };
                    c.track = function(c) {
                        var h = a.extend({},
                            l);
                        c || (c = {});
                        a.foreach(l,
                            function(d, n) {
                                a.exists(c[d]) && (h[d] = c[d], delete c[d])
                            });
                        return h
                    }
                }(jwplayer.playlist),
                function(c) {
                    function a(a, d, c) {
                        var b = a.style;
                        b.backgroundColor = "#000";
                        b.color = "#FFF";
                        b.width = l.styleDimension(c.width);
                        b.height = l.styleDimension(c.height);
                        b.display = "table";
                        b.opacity = 1;
                        c = document.createElement("p");
                        b = c.style;
                        b.verticalAlign = "middle";
                        b.textAlign = "center";
                        b.display = "table-cell";
                        b.font = "15px/20px Arial, Helvetica, sans-serif";
                        c.innerHTML = d.replace(":", ":<br>");
                        a.innerHTML = "";
                        a.appendChild(c)
                    }
                    var l = c.utils,
                        f = c.events,
                        h = document,
                        d = c.embed = function(n) {
                            function g(a,
                                b) {
                                l.foreach(b,
                                    function(b, e) {
                                        "function" == typeof a[b] && a[b].call(a, e)
                                    })
                            }

                            function m() {
                                if (!H)
                                    if ("array" == l.typeOf(k.playlist) && 2 > k.playlist.length && (0 == k.playlist.length || !k.playlist[0].sources || 0 == k.playlist[0].sources.length)) q();
                                    else if (!u)
                                    if ("string" == l.typeOf(k.playlist)) {
                                        var a = new c.playlist.loader;
                                        a.addEventListener(f.JWPLAYER_PLAYLIST_LOADED,
                                            function(a) {
                                                k.playlist = a.playlist;
                                                u = !1;
                                                m()
                                            });
                                        a.addEventListener(f.JWPLAYER_ERROR,
                                            function(a) {
                                                u = !1;
                                                q(a)
                                            });
                                        u = !0;
                                        a.load(k.playlist)
                                    } else if (z.getStatus() ==
                                    l.loaderstatus.COMPLETE) {
                                    for (a = 0; a < k.modes.length; a++)
                                        if (k.modes[a].type && d[k.modes[a].type]) {
                                            var h = l.extend({},
                                                    k),
                                                F = new d[k.modes[a].type](e, k.modes[a], h, z, n);
                                            if (F.supportsConfig()) return F.addEventListener(f.ERROR, b),
                                                F.embed(),
                                                g(n, h.events),
                                                n
                                        }
                                    var p;
                                    k.fallback ? (p = "No suitable players found and fallback enabled", A = setTimeout(function() {
                                            t(p, !0)
                                        },
                                        10), l.log(p), new d.download(e, k, q)) : (p = "No suitable players found and fallback disabled", t(p, !1), l.log(p), e.parentNode.replaceChild(E, e))
                                }
                            }

                            function b(a) {
                                p("\u64ad\u653e\u5668\u52a0\u8f7d\u9519\u8bef: " +
                                    a.message)
                            }

                            function r(a) {
                                p("Could not load plugins: " + a.message)
                            }

                            function q(a) {
                                a && a.message ? p("Error loading playlist: " + a.message) : p("\u64ad\u653e\u5668\u52a0\u8f7d\u9519\u8bef: \u89c6\u9891\u5730\u5740\u9519\u8bef")
                            }

                            function t(a, b) {
                                A && (clearTimeout(A), A = null);
                                A = setTimeout(function() {
                                        A = null;
                                        n.dispatchEvent(f.JWPLAYER_SETUP_ERROR, {
                                            message: a,
                                            fallback: b
                                        })
                                    },
                                    0)
                            }

                            function p(b) {
                                H || (k.fallback ? (H = !0, a(e, b, k), t(b, !0)) : t(b, !1))
                            }
                            var k = new d.config(n.config),
                                e,
                                y,
                                E,
                                v = k.width,
                                D = k.height,
                                z = c.plugins.loadPlugins(n.id,
                                    k.plugins),
                                u = !1,
                                H = !1,
                                A = null;
                            k.fallbackDiv && (E = k.fallbackDiv, delete k.fallbackDiv);
                            k.id = n.id;
                            y = h.getElementById(n.id);
                            k.aspectratio ? n.config.aspectratio = k.aspectratio : delete n.config.aspectratio;
                            e = h.createElement("div");
                            e.id = y.id;
                            e.style.width = 0 < v.toString().indexOf("%") ? v : v + "px";
                            e.style.height = 0 < D.toString().indexOf("%") ? D : D + "px";
                            y.parentNode.replaceChild(e, y);
                            this.embed = function() {
                                H || (z.addEventListener(f.COMPLETE, m), z.addEventListener(f.ERROR, r), z.load())
                            };
                            this.errorScreen = p;
                            return this
                        };
                    c.embed.errorScreen =
                        a
                }(jwplayer),
                function(c) {
                    function a(a) {
                        if (a.playlist)
                            for (var c = 0; c < a.playlist.length; c++) a.playlist[c] = new h(a.playlist[c]);
                        else {
                            var g = {};
                            f.foreach(h.defaults,
                                function(c, b) {
                                    l(a, g, c)
                                });
                            g.sources || (a.levels ? (g.sources = a.levels, delete a.levels) : (c = {},
                                l(a, c, "file"), l(a, c, "type"), g.sources = c.file ? [c] : []));
                            a.playlist = [new h(g)]
                        }
                    }

                    function l(a, c, g) {
                        f.exists(a[g]) && (c[g] = a[g], delete a[g])
                    }
                    var f = c.utils,
                        h = c.playlist.item;
                    (c.embed.config = function(d) {
                        var h = {
                            fallback: !0,
                            height: 270,
                            primary: "html5",
                            width: 480,
                            base: d.base ?
                                d.base : f.getScriptPath("jwplayer.js"),
                            aspectratio: ""
                        };
                        d = f.extend(h, c.defaults, d);
                        var h = {
                                type: "html5",
                                src: "http://open.yizhibo.tv/player/jwplayer.html5.js"
                            },
                            g = {
                                type: "flash",
                                src: "http://open.yizhibo.tv/player/jwplayer.flash.swf"
                            };
                        d.modes = "flash" == d.primary ? [g, h] : [h, g];
                        d.listbar && (d.playlistsize = d.listbar.size, d.playlistposition = d.listbar.position, d.playlistlayout = d.listbar.layout);
                        d.flashplayer && (g.src = d.flashplayer);
                        d.html5player && (h.src = d.html5player);
                        a(d);
                        g = d.aspectratio;
                        if ("string" == typeof g && f.exists(g)) {
                            var m =
                                g.indexOf(":"); - 1 == m ? h = 0 : (h = parseFloat(g.substr(0, m)), g = parseFloat(g.substr(m + 1)), h = 0 >= h || 0 >= g ? 0 : g / h * 100 + "%")
                        } else h = 0; - 1 == d.width.toString().indexOf("%") ? delete d.aspectratio : h ? d.aspectratio = h : delete d.aspectratio;
                        return d
                    }).addConfig = function(c, h) {
                        a(h);
                        return f.extend(c, h)
                    }
                }(jwplayer),
                function(c) {
                    var a = c.utils,
                        l = document;
                    c.embed.download = function(c, h, d) {
                        function n(b, e) {
                            for (var c = l.querySelectorAll(b), d = 0; d < c.length; d++) a.foreach(e,
                                function(a, b) {
                                    c[d].style[a] = b
                                })
                        }

                        function g() {
                            var b = "#" + c.id + " .jwdownload";
                            c.style.width = "";
                            c.style.height = "";
                            n(b + "display", {
                                width: a.styleDimension(Math.max(320, t)),
                                height: a.styleDimension(Math.max(180, p)),
                                background: "black center no-repeat " + (e ? "url(" + e + ")" : ""),
                                backgroundSize: "contain",
                                position: "relative",
                                border: "none",
                                display: "block"
                            });
                            n(b + "display div", {
                                position: "absolute",
                                width: "100%",
                                height: "100%"
                            });
                            n(b + "logo", {
                                top: y.margin + "px",
                                right: y.margin + "px",
                                background: "top right no-repeat url(" + y.prefix + y.file + ")"
                            });
                            n(b + "icon", {
                                background: "center no-repeat url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAgNJREFUeNrs28lqwkAYB/CZqNVDDj2r6FN41QeIy8Fe+gj6BL275Q08u9FbT8ZdwVfotSBYEPUkxFOoks4EKiJdaDuTjMn3wWBO0V/+sySR8SNSqVRKIR8qaXHkzlqS9jCfzzWcTCYp9hF5o+59sVjsiRzcegSckFzcjT+ruN80TeSlAjCAAXzdJSGPFXRpAAMYwACGZQkSdhG4WCzehMNhqV6vG6vVSrirKVEw66YoSqDb7cqlUilE8JjHd/y1MQefVzqdDmiaJpfLZWHgXMHn8F6vJ1cqlVAkEsGuAn83J4gAd2RZymQygX6/L1erVQt+9ZPWb+CDwcCC2zXGJaewl/DhcHhK3DVj+KfKZrMWvFarcYNLomAv4aPRSFZVlTlcSPA5fDweW/BoNIqFnKV53JvncjkLns/n/cLdS+92O7RYLLgsKfv9/t8XlDn4eDyiw+HA9Jyz2eyt0+kY2+3WFC5hluej0Ha7zQQq9PPwdDq1Et1sNsx/nFBgCqWJ8oAK1aUptNVqcYWewE4nahfU0YQnk4ntUEfGMIU2m01HoLaCKbTRaDgKtaVLk9tBYaBcE/6Artdr4RZ5TB6/dC+9iIe/WgAMYADDpAUJAxjAAAYwgGFZgoS/AtNNTF7Z2bL0BYPBV3Jw5xFwwWcYxgtBP5OkE8i9G7aWGOOCruvauwADALMLMEbKf4SdAAAAAElFTkSuQmCC)"
                            })
                        }

                        function m(a, b, e) {
                            a = l.createElement(a);
                            b && (a.className = "jwdownload" + b);
                            e && e.appendChild(a);
                            return a
                        }

                        function b(b) {
                            var e = m("iframe", "", c);
                            e.src = "http://www.youtube.com/embed/" + a.youTubeID(b);
                            e.width = t;
                            e.height = p;
                            e.style.border = "none"
                        }
                        var r = a.extend({},
                                h),
                            q,
                            t = r.width ? r.width : 480,
                            p = r.height ? r.height : 320,
                            k,
                            e,
                            y = h.logo ? h.logo : {
                                prefix: a.repo(),
                                file: "http://open.yizhibo.tv/player/logo.png",
                                margin: 10
                            };
                        (function() {
                            var h,
                                l,
                                n,
                                p;
                            p = r.playlist;
                            var y,
                                t,
                                u = ["mp4", "aac", "mp3"];
                            if (p && p.length) {
                                y = p[0];
                                t = y.sources;
                                for (p =
                                    0; p < t.length; p++) {
                                    var B = t[p],
                                        w = B.type ? B.type : a.extensionmap.extType(a.extension(B.file));
                                    B.file && a.foreach(u,
                                        function(b) {
                                            w == u[b] ? (h = B.file, l = y.image) : a.isYouTube(B.file) && (n = B.file)
                                        })
                                }
                                h ? (k = h, e = l, c && (q = m("a", "display", c), m("div", "icon", q), m("div", "logo", q), k && q.setAttribute("href", a.getAbsolutePath(k))), g()) : n ? b(n) : d()
                            }
                        })()
                    }
                }(jwplayer),
                function(c) {
                    var a = c.utils,
                        l = c.events,
                        f = {};
                    (c.embed.flash = function(d, n, g, m, b) {
                        function r(a, b, c) {
                            var d = document.createElement("param");
                            d.setAttribute("name", b);
                            d.setAttribute("value",
                                c);
                            a.appendChild(d)
                        }

                        function q(a, c, d) {
                            return function(f) {
                                try {
                                    d && document.getElementById(b.id + "_wrapper").appendChild(c);
                                    var k = document.getElementById(b.id).getPluginConfig("display");
                                    "function" == typeof a.resize && a.resize(k.width, k.height);
                                    c.style.left = k.x;
                                    c.style.top = k.h
                                } catch (g) {}
                            }
                        }

                        function t(b) {
                            if (!b) return {};
                            var c = {},
                                d = [];
                            a.foreach(b,
                                function(b, e) {
                                    var k = a.getPluginName(b);
                                    d.push(b);
                                    a.foreach(e,
                                        function(a, b) {
                                            c[k + "." + a] = b
                                        })
                                });
                            c.plugins = d.join(",");
                            return c
                        }
                        var p = new c.events.eventdispatcher,
                            k = a.flashVersion();
                        a.extend(this, p);
                        this.embed = function() {
                            g.id = b.id;
                            if (10 > k) return p.sendEvent(l.ERROR, {
                                message: "Flash version must be 10.0 or greater"
                            }), !1;
                            var e,
                                c,
                                h = b.config.listbar,
                                v = a.extend({},
                                    g);
                            if (d.id + "_wrapper" == d.parentNode.id) e = document.getElementById(d.id + "_wrapper");
                            else {
                                e = document.createElement("div");
                                c = document.createElement("div");
                                c.style.display = "none";
                                c.id = d.id + "_aspect";
                                e.id = d.id + "_wrapper";
                                e.style.position = "relative";
                                e.style.display = "block";
                                e.style.width = a.styleDimension(v.width);
                                e.style.height =
                                    a.styleDimension(v.height);
                                if (b.config.aspectratio) {
                                    var u = parseFloat(b.config.aspectratio);
                                    c.style.display = "block";
                                    c.style.marginTop = b.config.aspectratio;
                                    e.style.height = "auto";
                                    e.style.display = "inline-block";
                                    h && ("bottom" == h.position ? c.style.paddingBottom = h.size + "px" : "right" == h.position && (c.style.marginBottom = u / 100 * h.size * -1 + "px"))
                                }
                                d.parentNode.replaceChild(e, d);
                                e.appendChild(d);
                                e.appendChild(c)
                            }
                            e = m.setupPlugins(b, v, q);
                            0 < e.length ? a.extend(v, t(e.plugins)) : delete v.plugins;
                            "undefined" != typeof v["dock.position"] &&
                                "false" == v["dock.position"].toString().toLowerCase() && (v.dock = v["dock.position"], delete v["dock.position"]);
                            e = v.wmode ? v.wmode : v.height && 40 >= v.height ? "transparent" : "opaque";
                            c = "height width modes events primary base fallback volume".split(" ");
                            for (h = 0; h < c.length; h++) delete v[c[h]];
                            c = a.getCookies();
                            a.foreach(c,
                                function(a, b) {
                                    "undefined" == typeof v[a] && (v[a] = b)
                                });
                            c = window.location.href.split("/");
                            c.splice(c.length - 1, 1);
                            c = c.join("/");
                            v.base = c + "/";
                            f[d.id] = v;
                            a.isIE() ? (c = '<object classid="clsid:D27CDB6E-AE6D-11cf-96B8-444553540000" " width="100%" height="100%"id="' +
                                d.id + '" name="' + d.id + '" tabindex=0"">', c += '<param name="movie" value="' + n.src + '">', c += '<param name="allowfullscreen" value="true"><param name="allowscriptaccess" value="always">', c += '<param name="seamlesstabbing" value="true">', c += '<param name="wmode" value="' + e + '">', c += '<param name="bgcolor" value="#000000">', c += "</object>", d.outerHTML = c, e = document.getElementById(d.id)) : (c = document.createElement("object"), c.setAttribute("type", "application/x-shockwave-flash"), c.setAttribute("data", n.src), c.setAttribute("width",
                                "100%"), c.setAttribute("height", "100%"), c.setAttribute("bgcolor", "#000000"), c.setAttribute("id", d.id), c.setAttribute("name", d.id), c.setAttribute("tabindex", 0), r(c, "allowfullscreen", "true"), r(c, "allowscriptaccess", "always"), r(c, "seamlesstabbing", "true"), r(c, "wmode", e), d.parentNode.replaceChild(c, d), e = c);
                            b.config.aspectratio && (e.style.position = "absolute");
                            b.container = e;
                            b.setPlayer(e, "flash")
                        };
                        this.supportsConfig = function() {
                            if (k)
                                if (g) {
                                    if ("string" == a.typeOf(g.playlist)) return !0;
                                    try {
                                        var b = g.playlist[0].sources;
                                        if ("undefined" == typeof b) return !0;
                                        for (var c = 0; c < b.length; c++)
                                            if (b[c].file && h(b[c].file, b[c].type)) return !0
                                    } catch (d) {}
                                } else return !0;
                            return !1
                        }
                    }).getVars = function(a) {
                        return f[a]
                    };
                    var h = c.embed.flashCanPlay = function(c, f) {
                        if (a.isYouTube(c) || a.isRtmp(c, f) || "hls" == f) return !0;
                        var g = a.extensionmap[f ? f : a.extension(c)];
                        return g ? !!g.flash : !1
                    }
                }(jwplayer),
                function(c) {
                    var a = c.utils,
                        l = a.extensionmap,
                        f = c.events;
                    c.embed.html5 = function(h, d, n, g, m) {
                        function b(a, b, c) {
                            return function(d) {
                                try {
                                    var f = document.querySelector("#" +
                                        h.id + " .jwmain");
                                    c && f.appendChild(b);
                                    "function" == typeof a.resize && (a.resize(f.clientWidth, f.clientHeight), setTimeout(function() {
                                            a.resize(f.clientWidth, f.clientHeight)
                                        },
                                        400));
                                    b.left = f.style.left;
                                    b.top = f.style.top
                                } catch (g) {}
                            }
                        }

                        function r(a) {
                            q.sendEvent(a.type, {
                                message: "html5\u64ad\u653e\u5668\u52a0\u8f7d\u5931\u8d25"
                            })
                        }
                        var q = this,
                            t = new f.eventdispatcher;
                        a.extend(q, t);
                        q.embed = function() {
                            if (c.html5) {
                                g.setupPlugins(m, n, b);
                                h.innerHTML = "";
                                var l = c.utils.extend({},
                                    n);
                                delete l.volume;
                                l = new c.html5.player(l);
                                m.container = document.getElementById(m.id);
                                m.setPlayer(l, "html5")
                            } else l = new a.scriptloader(d.src),
                                l.addEventListener(f.ERROR, r),
                                l.addEventListener(f.COMPLETE, q.embed),
                                l.load()
                        };
                        q.supportsConfig = function() {
                            if (c.vid.canPlayType) try {
                                if ("string" == a.typeOf(n.playlist)) return !0;
                                for (var b = n.playlist[0].sources, d = 0; d < b.length; d++) {
                                    var e;
                                    var f = b[d].file,
                                        g = b[d].type;
                                    if (null !== navigator.userAgent.match(/BlackBerry/i) || a.isAndroid() && ("m3u" == a.extension(f) || "m3u8" == a.extension(f)) || a.isRtmp(f, g)) e = !1;
                                    else {
                                        var h =
                                            l[g ? g : a.extension(f)],
                                            m;
                                        if (!h || h.flash && !h.html5) m = !1;
                                        else {
                                            var q = h.html5,
                                                r = c.vid;
                                            if (q) try {
                                                m = r.canPlayType(q) ? !0 : !1
                                            } catch (t) {
                                                m = !1
                                            } else m = !0
                                        }
                                        e = m
                                    }
                                    if (e) return !0
                                }
                            } catch (t) {}
                            return !1
                        }
                    }
                }(jwplayer),
                function(c, a) {
                    var l = [],
                        f = c.utils,
                        h = c.events,
                        d = h.state,
                        n = document,
                        g = c.api = function(m) {
                            function b(a, b) {
                                return function(c) {
                                    return b(a, c)
                                }
                            }

                            function l(a, b) {
                                E[a] || (E[a] = [], p(h.JWPLAYER_PLAYER_STATE, q(a)));
                                E[a].push(b);
                                return e
                            }

                            function q(a) {
                                return function(b) {
                                    var c = b.newstate;
                                    b = b.oldstate;
                                    if (c == a) {
                                        var e = E[c];
                                        if (e)
                                            for (var d =
                                                    0; d < e.length; d++) {
                                                var f = e[d];
                                                "function" == typeof f && f.call(this, {
                                                    oldstate: b,
                                                    newstate: c
                                                })
                                            }
                                    }
                                }
                            }

                            function t(a, b) {
                                try {
                                    a.jwAddEventListener(b, 'function(dat) { jwplayer("' + e.id + '").dispatchEvent("' + b + '", dat); }')
                                } catch (c) {
                                    f.log("Could not add internal listener")
                                }
                            }

                            function p(a, b) {
                                u[a] || (u[a] = [], v && D && t(v, a));
                                u[a].push(b);
                                return e
                            }

                            function k() {
                                if (D) {
                                    if (v) {
                                        var a = Array.prototype.slice.call(arguments, 0),
                                            b = a.shift();
                                        if ("function" === typeof v[b]) {
                                            switch (a.length) {
                                                case 6:
                                                    return v[b](a[0], a[1], a[2], a[3], a[4], a[5]);
                                                case 5:
                                                    return v[b](a[0], a[1], a[2], a[3], a[4]);
                                                case 4:
                                                    return v[b](a[0], a[1], a[2], a[3]);
                                                case 3:
                                                    return v[b](a[0], a[1], a[2]);
                                                case 2:
                                                    return v[b](a[0], a[1]);
                                                case 1:
                                                    return v[b](a[0])
                                            }
                                            return v[b]()
                                        }
                                    }
                                    return null
                                }
                                z.push(arguments)
                            }
                            var e = this,
                                u = {},
                                E = {},
                                v,
                                D = !1,
                                z = [],
                                w,
                                x = {},
                                A = {};
                            e.container = m;
                            e.id = m.id;
                            e.getBuffer = function() {
                                return k("jwGetBuffer")
                            };
                            e.getContainer = function() {
                                return e.container
                            };
                            e.addButton = function(a, b, c, d) {
                                try {
                                    A[d] = c,
                                        k("jwDockAddButton", a, b, "jwplayer('" + e.id + "').callback('" + d + "')", d)
                                } catch (g) {
                                    f.log("Could not add dock button" +
                                        g.message)
                                }
                            };
                            e.removeButton = function(a) {
                                k("jwDockRemoveButton", a)
                            };
                            e.callback = function(a) {
                                if (A[a]) A[a]()
                            };
                            e.forceState = function(a) {
                                k("jwForceState", a);
                                return e
                            };
                            e.releaseState = function() {
                                return k("jwReleaseState")
                            };
                            e.getDuration = function() {
                                return k("jwGetDuration")
                            };
                            e.getFullscreen = function() {
                                return k("jwGetFullscreen")
                            };
                            e.getHeight = function() {
                                return k("jwGetHeight")
                            };
                            e.getLockState = function() {
                                return k("jwGetLockState")
                            };
                            e.getMeta = function() {
                                return e.getItemMeta()
                            };
                            e.getMute = function() {
                                return k("jwGetMute")
                            };
                            e.getPlaylist = function() {
                                var a = k("jwGetPlaylist");
                                "flash" == e.renderingMode && f.deepReplaceKeyName(a, ["__dot__", "__spc__", "__dsh__", "__default__"], [".", " ", "-", "default"]);
                                return a
                            };
                            e.getPlaylistItem = function(a) {
                                f.exists(a) || (a = e.getPlaylistIndex());
                                return e.getPlaylist()[a]
                            };
                            e.getPlaylistIndex = function() {
                                return k("jwGetPlaylistIndex")
                            };
                            e.getPosition = function() {
                                return k("jwGetPosition")
                            };
                            e.getRenderingMode = function() {
                                return e.renderingMode
                            };
                            e.getState = function() {
                                return k("jwGetState")
                            };
                            e.getVolume =
                                function() {
                                    return k("jwGetVolume")
                                };
                            e.getWidth = function() {
                                return k("jwGetWidth")
                            };
                            e.setFullscreen = function(a) {
                                f.exists(a) ? k("jwSetFullscreen", a) : k("jwSetFullscreen", !k("jwGetFullscreen"));
                                return e
                            };
                            e.setMute = function(a) {
                                f.exists(a) ? k("jwSetMute", a) : k("jwSetMute", !k("jwGetMute"));
                                return e
                            };
                            e.lock = function() {
                                return e
                            };
                            e.unlock = function() {
                                return e
                            };
                            e.load = function(a) {
                                k("jwLoad", a);
                                return e
                            };
                            e.playlistItem = function(a) {
                                k("jwPlaylistItem", parseInt(a, 10));
                                return e
                            };
                            e.playlistPrev = function() {
                                k("jwPlaylistPrev");
                                return e
                            };
                            e.playlistNext = function() {
                                k("jwPlaylistNext");
                                return e
                            };
                            e.resize = function(a, b) {
                                if ("flash" !== e.renderingMode) k("jwResize", a, b);
                                else {
                                    var c = n.getElementById(e.id + "_wrapper"),
                                        d = n.getElementById(e.id + "_aspect");
                                    d && (d.style.display = "none");
                                    c && (c.style.display = "block", c.style.width = f.styleDimension(a), c.style.height = f.styleDimension(b))
                                }
                                return e
                            };
                            e.play = function(b) {
                                b === a ? (b = e.getState(), b == d.PLAYING || b == d.BUFFERING ? k("jwPause") : k("jwPlay")) : k("jwPlay", b);
                                return e
                            };
                            e.pause = function(b) {
                                b === a ? (b =
                                    e.getState(), b == d.PLAYING || b == d.BUFFERING ? k("jwPause") : k("jwPlay")) : k("jwPause", b);
                                return e
                            };
                            e.stop = function() {
                                k("jwStop");
                                return e
                            };
                            e.seek = function(a) {
                                k("jwSeek", a);
                                return e
                            };
                            e.setVolume = function(a) {
                                k("jwSetVolume", a);
                                return e
                            };
                            e.createInstream = function() {
                                return new g.instream(this, v)
                            };
                            e.setInstream = function(a) {
                                return w = a
                            };
                            e.loadInstream = function(a, b) {
                                w = e.setInstream(e.createInstream()).init(b);
                                w.loadItem(a);
                                return w
                            };
                            e.getQualityLevels = function() {
                                return k("jwGetQualityLevels")
                            };
                            e.getCurrentQuality =
                                function() {
                                    return k("jwGetCurrentQuality")
                                };
                            e.setCurrentQuality = function(a) {
                                k("jwSetCurrentQuality", a)
                            };
                            e.getCaptionsList = function() {
                                return k("jwGetCaptionsList")
                            };
                            e.getCurrentCaptions = function() {
                                return k("jwGetCurrentCaptions")
                            };
                            e.setCurrentCaptions = function(a) {
                                k("jwSetCurrentCaptions", a)
                            };
                            e.getControls = function() {
                                return k("jwGetControls")
                            };
                            e.getSafeRegion = function() {
                                return k("jwGetSafeRegion")
                            };
                            e.setControls = function(a) {
                                k("jwSetControls", a)
                            };
                            e.destroyPlayer = function() {
                                k("jwPlayerDestroy")
                            };
                            e.playAd =
                                function(a) {
                                    var b = c(e.id).plugins;
                                    b.vast && b.vast.jwPlayAd(a)
                                };
                            e.pauseAd = function() {
                                var a = c(e.id).plugins;
                                a.vast ? a.vast.jwPauseAd() : k("jwPauseAd")
                            };
                            var B = {
                                onBufferChange: h.JWPLAYER_MEDIA_BUFFER,
                                onBufferFull: h.JWPLAYER_MEDIA_BUFFER_FULL,
                                onError: h.JWPLAYER_ERROR,
                                onSetupError: h.JWPLAYER_SETUP_ERROR,
                                onFullscreen: h.JWPLAYER_FULLSCREEN,
                                onMeta: h.JWPLAYER_MEDIA_META,
                                onMute: h.JWPLAYER_MEDIA_MUTE,
                                onPlaylist: h.JWPLAYER_PLAYLIST_LOADED,
                                onPlaylistItem: h.JWPLAYER_PLAYLIST_ITEM,
                                onPlaylistComplete: h.JWPLAYER_PLAYLIST_COMPLETE,
                                onReady: h.API_READY,
                                onResize: h.JWPLAYER_RESIZE,
                                onComplete: h.JWPLAYER_MEDIA_COMPLETE,
                                onSeek: h.JWPLAYER_MEDIA_SEEK,
                                onTime: h.JWPLAYER_MEDIA_TIME,
                                onVolume: h.JWPLAYER_MEDIA_VOLUME,
                                onBeforePlay: h.JWPLAYER_MEDIA_BEFOREPLAY,
                                onBeforeComplete: h.JWPLAYER_MEDIA_BEFORECOMPLETE,
                                onDisplayClick: h.JWPLAYER_DISPLAY_CLICK,
                                onControls: h.JWPLAYER_CONTROLS,
                                onQualityLevels: h.JWPLAYER_MEDIA_LEVELS,
                                onQualityChange: h.JWPLAYER_MEDIA_LEVEL_CHANGED,
                                onCaptionsList: h.JWPLAYER_CAPTIONS_LIST,
                                onCaptionsChange: h.JWPLAYER_CAPTIONS_CHANGED,
                                onAdError: h.JWPLAYER_AD_ERROR,
                                onAdClick: h.JWPLAYER_AD_CLICK,
                                onAdImpression: h.JWPLAYER_AD_IMPRESSION,
                                onAdTime: h.JWPLAYER_AD_TIME,
                                onAdComplete: h.JWPLAYER_AD_COMPLETE,
                                onAdCompanions: h.JWPLAYER_AD_COMPANIONS,
                                onAdSkipped: h.JWPLAYER_AD_SKIPPED
                            };
                            f.foreach(B,
                                function(a) {
                                    e[a] = b(B[a], p)
                                });
                            var C = {
                                onBuffer: d.BUFFERING,
                                onPause: d.PAUSED,
                                onPlay: d.PLAYING,
                                onIdle: d.IDLE
                            };
                            f.foreach(C,
                                function(a) {
                                    e[a] = b(C[a], l)
                                });
                            e.remove = function() {
                                if (!D) throw "Cannot call remove() before player is ready";
                                z = [];
                                g.destroyPlayer(this.id)
                            };
                            e.setup = function(a) {
                                if (c.embed) {
                                    var b = n.getElementById(e.id);
                                    b && (a.fallbackDiv = b);
                                    b = e;
                                    z = [];
                                    g.destroyPlayer(b.id);
                                    b = c(e.id);
                                    b.config = a;
                                    (new c.embed(b)).embed();
                                    return b
                                }
                                return e
                            };
                            e.registerPlugin = function(a, b, e, d) {
                                c.plugins.registerPlugin(a, b, e, d)
                            };
                            e.setPlayer = function(a, b) {
                                v = a;
                                e.renderingMode = b
                            };
                            e.detachMedia = function() {
                                if ("html5" == e.renderingMode) return k("jwDetachMedia")
                            };
                            e.attachMedia = function(a) {
                                if ("html5" == e.renderingMode) return k("jwAttachMedia", a)
                            };
                            e.removeEventListener = function(a, b) {
                                var c =
                                    u[a];
                                if (c)
                                    for (var e = c.length; e--;) c[e] === b && c.splice(e, 1)
                            };
                            e.dispatchEvent = function(a, b) {
                                var c = u[a];
                                if (c)
                                    for (var c = c.slice(0), e = f.translateEventResponse(a, b), d = 0; d < c.length; d++) {
                                        var k = c[d];
                                        if ("function" === typeof k) try {
                                            a === h.JWPLAYER_PLAYLIST_LOADED && f.deepReplaceKeyName(e.playlist, ["__dot__", "__spc__", "__dsh__", "__default__"], [".", " ", "-", "default"]),
                                                k.call(this, e)
                                        } catch (g) {
                                            f.log("There was an error calling back an event handler")
                                        }
                                    }
                            };
                            e.dispatchInstreamEvent = function(a) {
                                w && w.dispatchEvent(a, arguments)
                            };
                            e.callInternal = k;
                            e.playerReady = function(a) {
                                D = !0;
                                v || e.setPlayer(n.getElementById(a.id));
                                e.container = n.getElementById(e.id);
                                f.foreach(u,
                                    function(a) {
                                        t(v, a)
                                    });
                                p(h.JWPLAYER_PLAYLIST_ITEM,
                                    function() {
                                        x = {}
                                    });
                                p(h.JWPLAYER_MEDIA_META,
                                    function(a) {
                                        f.extend(x, a.metadata)
                                    });
                                for (e.dispatchEvent(h.API_READY); 0 < z.length;) k.apply(this, z.shift())
                            };
                            e.getItemMeta = function() {
                                return x
                            };
                            e.isBeforePlay = function() {
                                return k("jwIsBeforePlay")
                            };
                            e.isBeforeComplete = function() {
                                return k("jwIsBeforeComplete")
                            };
                            return e
                        };
                    g.selectPlayer =
                        function(a) {
                            var b;
                            f.exists(a) || (a = 0);
                            a.nodeType ? b = a : "string" == typeof a && (b = n.getElementById(a));
                            return b ? (a = g.playerById(b.id)) ? a : g.addPlayer(new g(b)) : "number" == typeof a ? l[a] : null
                        };
                    g.playerById = function(a) {
                        for (var b = 0; b < l.length; b++)
                            if (l[b].id == a) return l[b];
                        return null
                    };
                    g.addPlayer = function(a) {
                        for (var b = 0; b < l.length; b++)
                            if (l[b] == a) return a;
                        l.push(a);
                        return a
                    };
                    g.destroyPlayer = function(a) {
                        for (var b = -1, c, d = 0; d < l.length; d++) l[d].id == a && (b = d, c = l[d]);
                        0 <= b && (a = c.id, d = n.getElementById(a + ("flash" == c.renderingMode ?
                            "_wrapper" : "")), f.clearCss && f.clearCss("#" + a), d && ("html5" == c.renderingMode && c.destroyPlayer(), c = n.createElement("div"), c.id = a, d.parentNode.replaceChild(c, d)), l.splice(b, 1));
                        return null
                    };
                    c.playerReady = function(a) {
                        var b = c.api.playerById(a.id);
                        b ? b.playerReady(a) : c.api.selectPlayer(a.id).playerReady(a)
                    }
                }(window.jwplayer),
                function(c) {
                    var a = c.events,
                        l = c.utils,
                        f = a.state;
                    c.api.instream = function(c, d) {
                        function n(a, b) {
                            q[a] || (q[a] = [], d.jwInstreamAddEventListener(a, 'function(dat) { jwplayer("' + c.id + '").dispatchInstreamEvent("' +
                                a + '", dat); }'));
                            q[a].push(b);
                            return this
                        }

                        function g(b, c) {
                            t[b] || (t[b] = [], n(a.JWPLAYER_PLAYER_STATE, m(b)));
                            t[b].push(c);
                            return this
                        }

                        function m(a) {
                            return function(b) {
                                var c = b.newstate,
                                    d = b.oldstate;
                                if (c == a) {
                                    var f = t[c];
                                    if (f)
                                        for (var g = 0; g < f.length; g++) {
                                            var h = f[g];
                                            "function" == typeof h && h.call(this, {
                                                oldstate: d,
                                                newstate: c,
                                                type: b.type
                                            })
                                        }
                                }
                            }
                        }
                        var b,
                            r,
                            q = {},
                            t = {},
                            p = this;
                        p.type = "instream";
                        p.init = function() {
                            c.callInternal("jwInitInstream");
                            return p
                        };
                        p.loadItem = function(a, e) {
                            b = a;
                            r = e || {};
                            "array" == l.typeOf(a) ? c.callInternal("jwLoadArrayInstream",
                                b, r) : c.callInternal("jwLoadItemInstream", b, r)
                        };
                        p.removeEvents = function() {
                            q = t = {}
                        };
                        p.removeEventListener = function(a, b) {
                            var c = q[a];
                            if (c)
                                for (var d = c.length; d--;) c[d] === b && c.splice(d, 1)
                        };
                        p.dispatchEvent = function(a, b) {
                            var c = q[a];
                            if (c)
                                for (var c = c.slice(0), d = l.translateEventResponse(a, b[1]), f = 0; f < c.length; f++) {
                                    var g = c[f];
                                    "function" == typeof g && g.call(this, d)
                                }
                        };
                        p.onError = function(b) {
                            return n(a.JWPLAYER_ERROR, b)
                        };
                        p.onMediaError = function(b) {
                            return n(a.JWPLAYER_MEDIA_ERROR, b)
                        };
                        p.onFullscreen = function(b) {
                            return n(a.JWPLAYER_FULLSCREEN,
                                b)
                        };
                        p.onMeta = function(b) {
                            return n(a.JWPLAYER_MEDIA_META, b)
                        };
                        p.onMute = function(b) {
                            return n(a.JWPLAYER_MEDIA_MUTE, b)
                        };
                        p.onComplete = function(b) {
                            return n(a.JWPLAYER_MEDIA_COMPLETE, b)
                        };
                        p.onPlaylistComplete = function(b) {
                            return n(a.JWPLAYER_PLAYLIST_COMPLETE, b)
                        };
                        p.onPlaylistItem = function(b) {
                            return n(a.JWPLAYER_PLAYLIST_ITEM, b)
                        };
                        p.onTime = function(b) {
                            return n(a.JWPLAYER_MEDIA_TIME, b)
                        };
                        p.onBuffer = function(a) {
                            return g(f.BUFFERING, a)
                        };
                        p.onPause = function(a) {
                            return g(f.PAUSED, a)
                        };
                        p.onPlay = function(a) {
                            return g(f.PLAYING,
                                a)
                        };
                        p.onIdle = function(a) {
                            return g(f.IDLE, a)
                        };
                        p.onClick = function(b) {
                            return n(a.JWPLAYER_INSTREAM_CLICK, b)
                        };
                        p.onInstreamDestroyed = function(b) {
                            return n(a.JWPLAYER_INSTREAM_DESTROYED, b)
                        };
                        p.onAdSkipped = function(b) {
                            return n(a.JWPLAYER_AD_SKIPPED, b)
                        };
                        p.play = function(a) {
                            d.jwInstreamPlay(a)
                        };
                        p.pause = function(a) {
                            d.jwInstreamPause(a)
                        };
                        p.hide = function() {
                            c.callInternal("jwInstreamHide")
                        };
                        p.destroy = function() {
                            p.removeEvents();
                            c.callInternal("jwInstreamDestroy")
                        };
                        p.setText = function(a) {
                            d.jwInstreamSetText(a ? a : "")
                        };
                        p.getState = function() {
                            return d.jwInstreamState()
                        };
                        p.setClick = function(a) {
                            d.jwInstreamClick && d.jwInstreamClick(a)
                        }
                    }
                }(window.jwplayer)),
            jwplayer(I).setup({
                autostart: K,
                flashplayer: "http://open.yizhibo.tv/player/jwplayer.flash.swf",
                primary: "flash",
                file: J,
                image: C,
                width: w,
                height: x
            })
};