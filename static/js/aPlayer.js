!
function(k, s) {
    function u(a, c, b, d, f, e, g, h, m, n, v, r) {
        this.duration = parseInt(h) || 0;
        this.impression = f;
        this.clickUrl = c;
        this.tracking = b;
        this.event = d;
        this.oid = e;
        this.orderid = g;
        this.curIdx = v;
        this.resolveAdParam(a);
        this.adType = m + "";
        this.aduid = n;
        this.lc = r;
        this.initEvent()
    }
    function w(a) {
        this.support = "WebSocket" in window;
        this.ready = !1;
        this.target = a;
        this.mq = [];
        this.open()
    }
    var p = function() {
        var a = navigator.userAgent.toLowerCase();
        browser = {
            iPhone: /iphone/.test(a),
            iPad: /ipad/.test(a),
            iPod: /ipod/.test(a),
            isLetv: /letv/.test(a),
            Android: /android/.test(a),
            AndroidPad: /android/.test(a) && !/mobile/.test(a),
            atwin: /win/.test(a),
            opera: /opera/.test(a),
            msie: /msie/.test(a),
            firefox: /firefox/.test(a),
            safari: /safari/.test(a) && !/chrome/.test(a),
            wph: /windows phone/.test(a),
            ps: /playstation/.test(a),
            uc: /ucbrowser|ucweb/.test(a),
            xiaomi: /xiaomi/.test(a),
            weixin: /MicroMessenger/i.test(a),
            isLetvTv: function() {
                try {
                    return "function" == typeof LetvFish.getBrowserType
                } catch(a) {
                    return ! 1
                }
            }
        };
        var c = /(opera)(?:.*version)?[ \/]([\w.]+)/,
        b = /(msie) ([\w.]+)/,
        d = /(mozilla)(?:.*? rv:([\w.]+))?/,
        a = /(webkit)[ \/]([\w.]+)/.exec(a) || c.exec(a) || b.exec(a) || 0 > a.indexOf("compatible") && d.exec(a) || [];
        browser.version = a[2] || "0";
        return browser
    } (),
    l = function() {
        function a(a) {
            return c.call(a,
            function(a) {
                return null != a
            })
        }
        var c = [].filter,
        b = [].slice,
        d = /^\.([\w-]+)$/,
        f = /^#([\w-]*)$/,
        e = /^[\w-]+$/;
        vjs = function(a, b) {
            return new vjs.fn.init(a, b)
        };
        vjs.isPC = !1;
        var g = function(a, c) {
            var g;
            try {
                return null == a || 9 != a.nodeType && a.nodeType != a.DOCUMENT_NODE || !f.test(c) ? 1 !== a.nodeType && 9 !== a.nodeType ? [] : b.call(d.test(c) ? a.getElementsByClassName ? a.getElementsByClassName(RegExp.$1) : m(a, RegExp.$1) : e.test(c) ? a.getElementsByTagName(c) : a.querySelectorAll(c)) : (g = a.getElementById(RegExp.$1)) ? [g] : []
            } catch(h) {
                return []
            }
        },
        h = function(a, b, c) {
            b = b || [];
            a.selector = c || "";
            a.length = b.length;
            c = 0;
            for (var d = b.length; c < d; c++) a[c] = b[c];
            return a
        },
        m = function(a, b) {
            if (a.getElementsByTagName) for (var c = a.getElementsByTagName("*"), d = new RegExp("(^|\\s)" + b + "(\\s|$)"), e = 0, f = c.length; e < f; e++) if (d.test(c[e].className)) return [c[e]];
            return []
        };
        vjs.fn = {
            init: function(b, c) {
                if (b) {
                    if (b.nodeType) return h(this, [b]);
                    var d;
                    if (b instanceof Array) d = a(b);
                    else {
                        if (void 0 !== c) return vjs(c).find(b);
                        d = g(k, b)
                    }
                    return h(this, d, b)
                }
                return h(this)
            },
            find: function(a) {
                var b = this;
                return "object" == typeof a ? vjs(a).filter(function() {
                    var a = this;
                    return [].some.call(b,
                    function(b) {
                        return vjs.contains(b, a)
                    })
                }) : 1 == this.length ? vjs(g(this[0], a)) : this.map(function() {
                    return g(this, a)
                })
            },
            each: function(a) {
                if ([].every)[].every.call(this,
                function(b, c) {
                    return ! 1 !== a.call(b, c, b)
                });
                else for (var b = 0,
                c = this.length; b < c; b++) a.call(this[b], b, this[b]);
                return this
            },
            hasClass: function(a) {
                return (new RegExp("(\\s|^)" + a + "(\\s|$)")).test(this[0].className)
            },
            addClass: function(a) {
                var b = (a || "").split(/\s+/);
                return this.each(function() {
                    for (var a = this.className,
                    c = 0,
                    d = b.length; c < d; c++) vjs(this).hasClass(b[c]) || (a += " " + b[c]);
                    this.className = a
                })
            },
            removeClass: function(a) {
                var b = (a || "").split(/\s+/);
                return this.each(function() {
                    for (var a = this.className,
                    c = 0,
                    d = b.length; c < d; c++) a = a.replace(new RegExp("(\\s|^)" + b[c] + "(\\s|$)"), " ");
                    this.className = vjs.trim(a)
                })
            },
            on: function(a, b, c) {
                return this.each(function(d, e) {
                    var f = function(a) {
                        a.target = a.target || a.srcElement;
                        b.call(c, a)
                    };
                    e.domid || (e.domid = String(Math.random()).slice( - 4));
                    b[a + "_" + e.domid] = f;
                    e.addEventListener ? e.addEventListener(a, f, !1) : e.attachEvent && e.attachEvent("on" + a, f)
                })
            },
            off: function(a, b, c) {
                return this.each(function(c, d) {
                    var e = b[a + "_" + d.domid] || b;
                    d.removeEventListener ? d.removeEventListener(a, e, !1) : d.detachEvent && d.detachEvent("on" + a, e)
                })
            },
            getStyle: function(a) {
                var b = this[0];
                if (p.msie) {
                    switch (a) {
                    case "opacity":
                        return (b.filters["DXImageTransform.Microsoft.Alpha"] || b.filters.alpha || {}).opacity || 100;
                    case "float":
                        a = "styleFloat"
                    }
                    return b.style[a] || b.currentStyle ? b.currentStyle[a] : 0
                }
                "float" == a && (a = "cssFloat");
                return b.style[a] || (k.defaultView.getComputedStyle(b, "") ? k.defaultView.getComputedStyle(b, "")[a] : null) || 0
            },
            setStyle: function(a, b) {
                return this.each(function() {
                    if (p.msie) switch (a) {
                    case "opacity":
                        this.style.filter = "alpha(opacity=" + 100 * b + ")";
                        this.currentStyle && this.currentStyle.hasLayout || (this.style.zoom = 1);
                        return;
                    case "float":
                        a = "styleFloat"
                    } else "float" == a && (a = "cssFloat");
                    this.style[a] = b
                })
            },
            getAttr: function(a) {
                return this[0].getAttribute(a)
            },
            setAttr: function(a, b) {
                return this.each(function() {
                    this.setAttribute(a, b)
                })
            },
            offset: function() {
                var a = this[0],
                b = k.body,
                c = a.getBoundingClientRect();
                return {
                    top: c.top + (window.scrollY || b.parentNode.scrollTop || a.scrollTop) - (k.documentElement.clientTop || b.clientTop || 0),
                    left: c.left + (window.scrollX || b.parentNode.scrollLeft || a.scrollLeft) - (k.documentElement.clientLeft || b.clientLeft || 0)
                }
            },
            width: function(a) {
                if ("undefined" == typeof a) return this[0].offsetWidth;
                this[0].style.width = parseFloat(a) + "px"
            },
            height: function(a) {
                if ("undefined" == typeof a) return this[0].offsetHeight;
                this[0].style.height = parseFloat(a) + "px"
            },
            map: function(a) {
                return vjs(vjs.map(this,
                function(b, c) {
                    return a.call(b, c, b)
                }))
            }
        };
        vjs.fn.init.prototype = vjs.fn;
        vjs.contains = function(a, b) {
            return a !== b && a.contains(b)
        };
        vjs.map = function(a, b) {
            var c, d = [],
            e;
            if ("number" == typeof a.length) for (e = 0; e < a.length; e++) c = b(a[e], e),
            null != c && d.push(c);
            else for (e in a) c = b(a[e], e),
            null != c && d.push(c);
            return d
        };
        vjs.each = function(a, b) {
            var c;
            if ("number" == typeof a.length) for (c = 0; c < a.length && !1 !== b.call(this, c, a[c]); c++);
            else for (c in a) if (!1 === b.call(this, c, a[c])) break;
            return a
        };
        vjs.trim = function(a) {
            return a.replace(/^\s\s*/, "").replace(/\s\s*$/, "")
        };
        vjs.each("Boolean Number String Function Array Date RegExp Object".split(" "),
        function(a, b) {
            b.toLowerCase()
        });
        vjs.extend = function(a) {
            var b, c;
            a = a || {};
            b = a.init || a.init || this.prototype.init || this.prototype.init ||
            function() {};
            c = function() {
                b.apply(this, arguments)
            };
            c.prototype = vjs.create(this.prototype);
            c.prototype.constructor = c;
            a.__proto__ || (a.__proto__ = c.prototype);
            c.extend = vjs.extend;
            for (var d in a) a.hasOwnProperty(d) && (c.prototype[d] = a[d]);
            return c
        };
        vjs.create = function(a) {
            function b() {}
            b.prototype = a;
            return new b
        };
        vjs.getWinWH = function() {
            var a = window.innerWidth,
            b = window.innerHeight;
            "number" != typeof a && ("CSS1Compat" == k.compatMode ? (a = k.documentElement.clientWidth, b = k.documentElement.clientHeight) : (a = k.body.clientWidth, b = k.body.clientHeight));
            return {
                width: a,
                height: b
            }
        };
        vjs.safari = p.safari;
        return vjs
    } (),
    q = {
        createElement: function(a, c) {
            var b = k.createElement(a),
            d;
            for (d in c) c.hasOwnProperty(d) && ( - 1 !== d.indexOf("-") ? b.setAttribute(d, c[d]) : b[d] = c[d]);
            return b
        },
        now: Date.now ||
        function() {
            return + new Date
        },
        merge: function(a, c, b) {
            c || (c = {});
            for (var d in c) ! c.hasOwnProperty(d) || b && a.hasOwnProperty(d) || (a[d] = c[d]);
            return a
        },
        getJSON: function(a, c, b, d, f) {
            var e = this.now(),
            g = "vjs_" + e + Math.floor(100 * Math.random()),
            h = "$1" + g + "$2",
            m = 0,
            n = 0,
            v = this,
            r,
            l = k.head || k.getElementsByTagName("head")[0] || k.documentElement;
            /_r=/i.test(a) || (a += "&_r=?");
            a = a.replace(/(\=)\?(&|$)|\?\?/i, h);
            d = d || 5E3;
            f = f || 2;
            window[g] = function(a) {
                q();
                vjs.responseTime = v.now() - e;
                vjs.retryCount = m;
                c.call(this, a, {
                    responseTime: v.now() - e,
                    retryCount: m
                });
                window[g] = null
            };
            var q = function() {
                clearTimeout(n);
                r && r.parentNode && (l.removeChild(r), r.onload = r.onreadystatechange = null, r = s)
            },
            p = function() {
                q();
                m >= f ? (clearTimeout(n), window[g] = null, b && b.call(this)) : (a = a.replace(/&_r=[\d|\?]+/i, "&_r=" + m), r = k.createElement("script"), r.setAttribute("type", "text/javascript"), r.setAttribute("src", a), r.onload = r.onreadystatechange = function(a) {
                    clearTimeout(n)
                },
                l.insertBefore(r, l.firstChild), n = setTimeout(p, d), m++)
            };
            p()
        }
    },
    y = function() {
        var a = screen;
        return {
            x: a.width > a.height ? a.width: a.height,
            y: a.width > a.height ? a.height: a.width
        }
    } (),
    x;
    x = p.iPhone || p.iPod || p.msie ? !1 : !0;
    var g = {
        param: function(a) {
            var c = [];
            if ("object" == typeof a) for (var b in a) a.hasOwnProperty(b) && "-" !== a[b] && c.push(encodeURIComponent(b) + "=" + encodeURIComponent(a[b]));
            return c.join("&")
        },
        sendLogs: function(a, c) {
            if (g.getQuery("arkdebug")) try {
                var b = this.el("#arkDebugButton"),
                d = h.putinVars.uuid,
                f = "http://ark.letv.com/apsdbg/?type=1&sid=" + d + "&time=" + c + "&msg=";
                g.existEl(b) || (b = q.createElement("div", {
                    id: "arkDebugButton",
                    className: "vdo_send_log"
                }), b.innerText = "\u8bf7\u67e5\u770b" + d + "\u7684\u65e5\u5fd7", h.staticVars.countdownElem.appendChild(b)); (new Image).src = f + encodeURIComponent("`" + a + "`")
            } catch(e) {}
        },
        wsLog: function(a) {
            if ("2" == g.getQuery("arkdebug")) try {
                this.s || (this.s = new w(h.config.WS_URL)),
                this.s.addLog(a)
            } catch(c) {
                console.log(c)
            }
        },
        debug: function(a, c, b) {
            b = b || " ";
            if (!0 == h.config.DEBUG || g.getQuery("arkdebug"))"object" == typeof a ? (c && (console.log("%c" + c, "color:#f0d"), this.wsLog(c)), this.wsLog(a), console.dir(a)) : a == s ? console.log("\u6570\u636e\u7a7a" + b) : (this.wsLog(a), console.log(a + b))
        },
        json: function(a) {
            return "string" === typeof a ? JSON && JSON.parse ? JSON.parse(a) : eval("(" + a + ")") : JSON.stringify(a)
        },
        resoSid: function(a) {
            var c = "";
            "string" === typeof a && (a = a.split(",")[0], a = a.split("|"), c = 3 == a.length ? a[1] : 1 < a.length ? a[1] : a[0]);
            return c
        },
        getQuery: function(a, c) {
            var b = c || location.search;
            return 0 < b.length && -1 != b.indexOf("?") ? (b = b.match(new RegExp(a + "=([^&]*)", "i"))) && 0 < b.length ? unescape(b[1]) : null: null
        },
        easyClone: function(a, c) {
            for (var b in c) c.hasOwnProperty(b) && "object" !== typeof c[b] && (a[b] = c[b])
        },
        arkMapper: function(a) {
            if ("string" == typeof a && (a = parseInt(a), isNaN(a))) return 92;
            a = h.config.ARK_Mapper[a] || a;
            if (this.isMStation && (p.iPhone || p.iPod)) {
                var c = h.config.M_ARK_MAPPER[a];
                if (c) return c
            }
            return a
        },
        removeElem: function(a) {
            if (a) return a.remove ? a.remove() : a.parentNode && a.parentNode.removeChild && a.parentNode.removeChild(a)
        },
        el: function(a, c) {
            var b = c ? l(a).find(c)[0] : l(a)[0];
            b || (b = {
                setAttribute: function() {},
                style: {},
                isnull: !0
            });
            return b
        },
        existEl: function(a) {
            return a && !a.isnull ? a instanceof Array ? 0 < a.length: !0 : !1
        },
        getAslbUrl: function(a, c) {
            var b, d;
            c.result = c.result || [];
            if (a instanceof Array) {
                if (d = a.shift()) 0 <= d.url.indexOf(h.config.ASLB_DOMAIN) ? b = p.iPhone || p.iPod || p.iPad ? d.url + "&tss=ios&format=1&jsonp=?": d.url + "&format=1&jsonp=?": (c.result.push(d), g.getAslbUrl(a, c));
                else return c(c.result);
                b !== s && q.getJSON(b,
                function(b) {
                    if (!1 == /mp4|m3u8/.test(b.location)) {
                        debugger;
                        d.ryCount = l.retryCount;
                        d.costTime = l.responseTime;
                        d.err = 474;
                        h.sendEvent(h.config.SEND_EVENT_TYPE.OnASLB, {
                            curAD: d,
                            curIndex: d.curIdx
                        });
                        h.collectError("474,format error," + g.json(d), 3)
                    } else d.rUrl = d.url,
                    d.url = b.location,
                    d.ryCount = l.retryCount,
                    d.costTime = l.responseTime,
                    c.result.push(d),
                    h.sendEvent(h.config.SEND_EVENT_TYPE.OnASLB, {
                        curAD: d,
                        curIndex: d.curIdx
                    });
                    g.getAslbUrl(a, c)
                },
                function(b) {
                    d.ryCount = l.retryCount;
                    d.costTime = l.responseTime;
                    d.err = 473;
                    c.result.push(d);
                    h.sendEvent(h.config.SEND_EVENT_TYPE.OnASLB, {
                        curAD: d,
                        curIndex: d.curIdx
                    });
                    h.collectError("473,aslb error," + g.json(d), 3);
                    g.getAslbUrl(a, c)
                })
            } else {
                if (0 > b.indexOf(h.config.ASLB_DOMAIN)) return c([]);
                q.getJSON(a + "&format=1&jsonp=?",
                function(a) {
                    return c([a.location])
                },
                function(a) {
                    return c([])
                })
            }
        },
        loadCss: function(a) {
            var c = k.head || k.getElementsByTagName("head")[0] || k.documentElement,
            b = k.createElement("style");
            b.setAttribute("type", "text/css");
            b.innerHTML = a;
            c.appendChild(b)
        },
        getDeviceSize: y,
        canBeClicked: x,
        isUC: p.uc,
        isMStation: !1
    },
    h = {
        dynamicVars: {
            retry: 0,
            adidQueue: [],
            isFirst: !0,
            hasPlayed: !1
        },
        staticVars: {
            arkId: 92,
            countdownElem: null
        },
        putinVars: {},
        config: {
            AD_STYLE: {
                pre_roll: "2",
                standard: "3",
                pause: "6"
            },
            SEND_EVENT_TYPE: {
                OnStart: "AD_PLAY",
                OnComplate: "AD_ENDED",
                OnClick: "AD_CLICK",
                OnAcComplate: "AC_COMPLATE",
                OnError: "AD_ERROR",
                OnPause: "AD_PAUSE",
                OnASLB: "AD_ASLB",
                OnLoginAc: "loginCb"
            },
            CALL_PLAYER_TYPE: {
                playAD: "playAD",
                stopAD: "stopAD",
                pauseAD: "pauseAD",
                resumeAD: "resumeAD",
                getRealTime: "getCurrTime",
                getPlayerSize: "getVideoRect",
                doLogin: "login",
                pingback: "pingback"
            },
            PROCESS_EVENT_TICKS: [{
                k: "firstQuartile",
                v: 0.25
            },
            {
                k: "midpoint",
                v: 0.5
            },
            {
                k: "thirdQuartile",
                v: 0.75
            }],
            crc_table: [61888, 62024, 21822, 44648, 51027, 25193, 39449, 32749, 45072, 19780, 27911, 40640, 22412, 47959, 2033, 15647, 26948, 7977, 333, 52810, 2229, 28457, 56115, 3222, 7819, 8261, 37040, 26479, 46017, 37654, 52255, 36436, 49642, 26018, 41611, 57969, 22529, 40087, 25454, 12785, 50531, 1739, 4421, 44187, 14573, 60124, 48843, 50551, 63571, 18928, 9702, 31935, 37924, 53689, 43138, 29106, 22299, 17913, 22765, 17733, 13233, 54102, 63095, 54790, 45315, 4283, 52320, 21487, 24719, 23499, 25688, 43296, 18522, 46226, 54051, 23750, 63855, 40050, 23830, 13909, 53473, 35269, 6541, 59749, 45495, 7225, 26512, 17657, 28777, 4159, 17208, 50565, 48334, 33575, 10897, 26141, 42425, 51911, 4632, 28267, 27030, 57778, 15356, 31158, 14774, 53522, 27342, 33231, 29241, 52365, 12102, 5400, 40637, 7989, 51774, 31639, 1064, 46043, 38691, 42315, 25171, 2606, 94, 25879, 50273, 48389, 61059, 63334, 38144, 34805, 17489, 9758, 21488, 31104, 40127, 47832, 19575, 8379, 62899, 64770, 6327, 15962, 35087, 34E3, 41978, 50244, 40758, 57390, 20080, 51537, 61759, 31722, 57084, 25726, 3693, 42772, 41971, 46086, 30626, 46885, 37383, 847, 38119, 23229, 59572, 58742, 40006, 20034, 62943, 57283, 50816, 54485, 36496, 28963, 5481, 23375, 51432, 3135, 18675, 20557, 968, 55963, 47914, 45119, 25284, 1646, 34994, 1493, 10573, 32670, 64131, 45013, 56896, 57534, 26361, 47505, 26941, 31536, 886, 43364, 32112, 18014, 13600, 60378, 12717, 60596, 9862, 56041, 44055, 39986, 37168, 28168, 55209, 30733, 5480, 6034, 17485, 56710, 63417, 33557, 9848, 39651, 64250, 14639, 63835, 38963, 7906, 39909, 7971, 10158, 40564, 25844, 3305, 50258, 28353, 42316, 44088, 44477, 1500, 42481, 45659, 44289, 10989, 54239, 19915, 42407, 19391, 1463, 50295, 60742, 8528, 50215, 445, 89, 39965, 42071],
            ARK_Mapper: {
                106 : "531"
            },
            M_ARK_MAPPER: {
                531 : "531"
            },
            H5_ADPLAYER_VER: "aps_h5_2.0.13",
            COUNTDOWN_IMG_URL: "http://i2.letvimg.com/img/201310/21/numbers.png",
            ARK_DOMAIN: "ark.letv.com",
            ASLB_DOMAIN: "g3.letv",
            ARK_SHOW_URL: "http://ark.letv.com/s?res=jsonp",
            ARK_PREVIEW_URL: "http://ark.letv.com/p?res=jsonp",
            DC_AD_URL: "http://dc.letv.com/va/?",
            SKIP_AD_CLICK: "http://dc.letv.com/s/?k=sumtmp;H5PADQad",
            SKIP_AD_SUCC: "http://dc.letv.com/s/?k=sumtmp;H5PADQadfc",
            REQ_ARK_TIMEOUT: 5E3,
            DOWNLOAD_URL_TIMEOUT: 1E4,
            WS_URL: "ws://10.58.88.69:8080",
            CSS_TEMPLATE: ".aps_countdown_cont{position:absolute;border-radius:10px 0;top:10px;right:10px;display:block;padding:5px 10px;background:rgba(49,37,37,0.8);z-index:12} .precdImg{float:left;width:12px;height:20px;overflow:hidden;}.vdo_post_time,.vdo_post_detail{position:absolute;height:40px;border:1px solid #262626;text-align:center;line-height:40px;font-size:16px;z-index:13;}.vdo_post_time{right:40px;top:20px;color:#ccc;}.vdo_post_rlt{position:relative;width:100%;height:40px}.vdo_time_bg,.vdo_detail_bg{position:absolute;width:100%;height:40px;left:0;top:0;background-color:#000;opacity:0.7}.vdo_time_info,.vdo_detail_info{padding:0 10px;position:relative}.vdo_detail_info{padding:0 20px}.vdo_time_info span{color:#09adfe;padding:9px 5px 0 0;float:right}.vdo_time_info a{color:#cccccc;margin-left:3px;text-decoration: none;}.vdo_post_detail{left:40px;bottom:20px}.vdo_post_detail a,.vdo_post_detail a:hover{color:#ccc;display:block;width:100%;height:40px;text-decoration: none;}.vdo_post_detail i{background:url(http://i3.letvimg.com/img/201404/15/1052/rightLink.png) no-repeat left top;width:14px;height:24px;float:right;margin:8px 0 0 10px}.hv_box_mb .vdo_post_time,.hv_box_live_mb .vdo_post_time{right:10px;top:10px;}.hv_box_mb .vdo_post_detail,.hv_box_live_mb .vdo_post_detail{left:10px;bottom:10px}.hv_box_mb .vdo_post_time,.hv_box_mb .vdo_post_detail,.hv_box_live_mb .vdo_post_time,.hv_box_live_mb .vdo_post_detail{height:30px;line-height:30px;font-size:13px}.hv_box_mb .vdo_post_rlt,.hv_box_mb .vdo_time_bg,.hv_box_mb .vdo_detail_bg,.hv_box_mb .vdo_post_detail a,.hv_box_mb .vdo_post_detail,.hv_box_live_mb .vdo_post_rlt,.hv_box_live_mb .vdo_time_bg,.hv_box_live_mb .vdo_detail_bg,.hv_box_live_mb .vdo_post_detail a,.hv_box_live_mb .vdo_post_detail a:hover{height:30px}.hv_box_mb .vdo_detail_info,.hv_box_live_mb .vdo_detail_info{padding:0 10px}.hv_box_mb .vdo_post_detail i,.hv_box_live_mb .vdo_post_detail i{width:7px;height:12px;background-size:100%;margin:8px 0 0 5px}.hv_box_mb .vdo_time_info span,.hv_box_live_mb .vdo_time_info span{color:#09adfe;padding:4px 0px 0 0;float:right}.aps_mask_cont{position: absolute;width: 100%;height: 100%;top: 0px;left: 0px;z-index: 12;}.aps_pop_poster{width:100%;height:100%;position:absolute;top:0;left:0;z-index:150;}.vdo_send_log{position:absolute;top:80px;height:100px;right:10px;font-size:30px;z-index:30}.hv_pop_poster{position:absolute;top:50%;left:50%;margin:-112px 0 0 -182px; width:365px;height:225px;overflow:hidden;background-color:#f1f1f1;}.hv_pop_poster p{text-align:center;margin-bottom:12px}.hv_pop_poster p.hv_p1{padding-top:48px}.hv_pop_poster a{display:inline-block;height:40px;width:224px;line-height:40px;background-color:#f7f7f7;font-size:15px;color:#7e7e7e;border:1px solid #d1d1d1}.hv_pop_poster a.blu{background-color:#00a0e9;color:#ffffff;border:1px solid #00a0e9}.hv_pop_poster a.close{width:20px;height:20px;display:block;position:absolute;top:10px;right:10px;border:none;background:none}.hv_pop_poster a.close i{display:block;width:18px;height:2px;position:absolute;top:6px;left:0;background:#737373;transform:rotate(-45deg);-ms-transform:rotate(-45deg);   -moz-transform:rotate(-45deg);  -webkit-transform:rotate(-45deg);-o-transform:rotate(-45deg)}.hv_pop_poster a.close i.i_1{transform:rotate(45deg);-ms-transform:rotate(45deg);  -moz-transform:rotate(45deg);   -webkit-transform:rotate(45deg);-o-transform:rotate(45deg)}.hv_pop_poster .hv_org{color:#fd6c01}.hv_ico_pasued, .hv_ico_loading {position: absolute;top: 50%;left: 50%;margin: -55px 0 0 -55px;width: 110px;height: 110px;overflow: hidden;z-index: 100;background: url(http://i3.letvimg.com/img/201403/24/hv_ico.png) no-repeat -140px 0;}",
            DEBUG: !1,
            ArkDebug: !1
        },
        adQueue: [],
        loadCss: function() {
            g.loadCss(this.config.CSS_TEMPLATE)
        },
        prepareImages: function(a, c) {
            var b = new Image;
            b.src = a;
            "undefined" != typeof c && (b.complete ? c(b.width, b.height) : b.onload = function() {
                c(b.width, b.height);
                b.onload = null
            })
        },
        destory: function(a) {
            try {
                a.closeCountDown(),
                this.callback2Player = null,
                this.putinVars = {},
                this.dynamicVars = {
                    retry: 0,
                    adidQueue: [],
                    isFirst: !0,
                    hasPlayed: !1
                },
                this.adQueue = [],
                this.playingMonitorCount = 0,
                this.playAdTimer && 0 < this.playAdTimer.length && clearTimeout(this.playAdTimer[a.curIndex]),
                clearTimeout(this.downMaterialTimer),
                clearTimeout(this.arkTimer),
                clearTimeout(this.playingMonitor),
                clearInterval(a.processTimer),
                clearInterval(a.countdownTimer)
            } catch(c) {}
        },
        openApp: function(a, c) {
            var b = "letvclient://msiteAction?actionType=0&pid=" + encodeURIComponent(a) + "&vid=" + encodeURIComponent(c) + "&from=mletv";
            setTimeout(function() {
                if (p.Android) {
                    var a = k.createElement("iframe");
                    a.style.cssText = "width:0px;height:0px;position:fixed;top:0;left:0;";
                    a.src = b;
                    k.body.appendChild(a)
                } else location.href = b;
                setTimeout(function() {},
                1500)
            },
            100)
        },
        initAD: function(a, c) {
            var b = this,
            d = "";
            u.curAd && b.destory(u.curAd);
            b.loadCss();
            b.prepareImages(b.config.COUNTDOWN_IMG_URL);
            g.debug(a, "\u4f20\u8fc7\u6765\u7684\u503c\uff1a");
            var f = b.config,
            e = f.SEND_EVENT_TYPE,
            t = f.CALL_PLAYER_TYPE;
            a && c ? (b.callback2Player = function() {
                try {
                    return c.apply(b, arguments)
                } catch(a) {
                    b.collectError("497&err=" + (a || {}).stack, 3)
                }
            },
            b.putinVars = a) : (a = b.putinVars, c = b.callback2Player);
            b.putinVars && b.putinVars.gdur && (d = parseInt(b.putinVars.gdur) || 0, d = 60 >= d ? 19999997 : 300 >= d ? 19999998 : 19999999);
            if (p.isLetv) b.sendEvent(e.OnAcComplate, {
                atype: "2",
                curAD: {},
                curIndex: -1,
                ia: 0
            }),
            b.callback2Player.call(b, t.stopAD);
            else if (b.startTime = q.now(), a.isvip) b.callback2Player.call(b, t.stopAD),
            b.sendEvent(e.OnAcComplate, {
                atype: "2",
                curAD: {},
                curIndex: -1,
                ia: a.isvip
            }),
            "MPlayer" != a.pname && b.sendEvent(e.OnAcComplate, {
                atype: "3",
                curAD: {},
                curIndex: -1,
                ia: a.isvip
            }),
            b.tips("tips", "\u60a8\u6b63\u4eab\u53d7\u4e50\u89c6\u4f1a\u5458\u53bb\u5e7f\u544a\u670d\u52a1");
            else if (a.isTrylook) b.callback2Player.call(b, t.stopAD),
            b.sendEvent(e.OnAcComplate, {
                atype: "2",
                curAD: {},
                curIndex: -1,
                ia: 3
            }),
            "MPlayer" != a.pname && b.sendEvent(e.OnAcComplate, {
                atype: "3",
                curAD: {},
                curIndex: -1,
                ia: 3
            }),
            b.tips("tips", "\u8bd5\u770b\u670d\u52a1");
            else if ("baidullq" == g.getQuery("ref") && 0 <= navigator.userAgent.indexOf("baidubrowser")) b.callback2Player.call(b, t.stopAD),
            b.sendEvent(e.OnAcComplate, {
                atype: "2",
                curAD: {},
                curIndex: -1,
                ia: 10
            }),
            b.tips("tips", "\u767e\u5ea6\u6e20\u9053\u7981\u64ad");
            else {
                "MPlayer" == a.pname ? (g.isMStation = !0, b.adStyle = a.style || f.AD_STYLE.pre_roll) : b.adStyle = a.style || [f.AD_STYLE.pre_roll, f.AD_STYLE.standard];
                try {
                    b.staticVars.countdownElem = g.el("#" + b.putinVars.cont, "div")
                } catch(k) {
                    h.callback2Player.call(b, t.stopAD, []);
                    b.sendEvent(e.OnAcComplate, {
                        error: {
                            code: 22
                        }
                    });
                    return
                }
                t = b.putinVars.ark ? b.putinVars.ark: "__ADINFO__" in window && __ADINFO__.arkId ? __ADINFO__.arkId: b.putinVars.streamid ? "!": "92";
                b.staticVars.arkId = g.arkMapper(t);
                b.arkTimer = setTimeout(function() {
                    g.debug("\u8bf7\u6c42ark\u8d85\u65f6,\u64ad\u653e\u6b63\u7247");
                    b.sendEvent(e.OnAcComplate, {
                        error: {
                            code: 451
                        }
                    });
                    h.callback2Player.call(h, f.CALL_PLAYER_TYPE.stopAD, [])
                },
                f.REQ_ARK_TIMEOUT);
                b.getArkData(b.adStyle, b.staticVars.arkId, d, b.putinVars.streamid)
            }
        },
        getArkData: function(a, c, b, d) {
            var f = this,
            e = f.config,
            h = e.SEND_EVENT_TYPE,
            k = f.dynamicVars,
            m;
            a instanceof Array && (a = a.join(","));
            b = {
                ark: c,
                n: k.isFirst ? 1 : 0,
                ct: a,
                vid: b || 0
            };
            "undefined" != typeof d && (d = g.resoSid(d), m = g.isMStation ? p.iPhone || p.iPod ? "471": "335": "148", g.easyClone(b, {
                sid: d,
                vid: "19999999",
                b: "2",
                ark: m
            }), f.staticVars.arkId = m);
            if (d = {
                coop_yinliu: 393,
                coop_yinliu1: 394,
                coop_yinliu2: 395,
                coop_yinliu3: 396
            } [g.getQuery("q2")]) b.ark = f.staticVars.arkId = d;
            var n = [this.config.ARK_SHOW_URL, g.param(b), "j=?"].join("&");
            d = {
                r: g.getQuery("r"),
                o: g.getQuery("o"),
                d: g.getQuery("d"),
                w: g.getQuery("w"),
                x: g.getQuery("x"),
                y: g.getQuery("y"),
                z: g.getQuery("z")
            };
            k.isFirst = !1;
            d.w && d.x && d.y && d.z && (n = [this.config.ARK_PREVIEW_URL, g.param(b), g.param(d), "j=?"].join("&"));
            d = null;
            g.debug("\u8bf7\u6c42ARK\u5730\u5740:" + n);
            q.getJSON(n,
            function(b) {
                try {
                    f._resolveData.call(f, b, a, n, c)
                } catch(d) {
                    f.callback2Player(e.CALL_PLAYER_TYPE.playAD, []),
                    f.sendEvent(h.OnAcComplate, {
                        error: {
                            code: 453
                        }
                    }),
                    g.debug(d, "\u89e3\u6790\u5f02\u5e38\uff1a")
                }
                clearTimeout(f.arkTimer)
            },
            function(a) {
                f.sendEvent(h.OnAcComplate, {
                    error: {
                        code: 450
                    }
                });
                f.callback2Player(e.CALL_PLAYER_TYPE.stopAD, []);
                clearTimeout(f.arkTimer)
            },
            e.REQ_ARK_TIMEOUT)
        },
        tips: function(a, c, b) {
            switch (a) {
            case "tips":
                g.debug(c)
            }
        },
        _resolveData: function(a, c, b, d) {
            var f = this,
            e = f.config;
            c = "-";
            if (a && a.vast) {
                b = a.vast;
                a = b.Ad.length;
                g.easyClone(f.staticVars, b);
                f.dynamicVars.preAdCount = 0;
                f.dynamicVars.staAdCount = 0;
                g.debug("\u8fd4\u56de\u5e7f\u544a\u6570\uff1a" + a);
                f.adQueue = [];
                f.dynamicVars.dur_total = 0;
                f.dynamicVars.dur = [];
                for (d = 0; d < a; d++) {
                    var h = b.Ad[d],
                    k = h.InLine,
                    m = h.cuepoint_type,
                    n = k.Creatives.Creative[0],
                    l = {};
                    1 === a && this.adStyle instanceof Array && (m == e.AD_STYLE.pre_roll ? this.adStyle.pop() : m == e.AD_STYLE.standard && this.adStyle.shift());
                    g.easyClone(l, h);
                    h = new u(n.Linear.AdParameters, n.Linear.VideoClicks.ClickThrough, n.Linear.VideoClicks.ClickTracking, n.Linear.TrackingEvents.Tracking, k.Impression, l.order_item_id, l.order_id, n.Linear.Duration, m, n.Linear.adzone_id, d, h.lc);
                    f.adQueue.push(h);
                    l.duration = h.duration;
                    m == e.AD_STYLE.pre_roll ? (f.dynamicVars.dur.push(l.duration), f.dynamicVars.dur_total += l.duration, f.dynamicVars.preAdCount++, f.dynamicVars.adidQueue.push(l.order_item_id)) : m == e.AD_STYLE.standard && (f.dynamicVars.staAdCount++, f.dynamicVars.stadur = l.duration, c = l.order_item_id)
                }
                g.getAslbUrl(f.adQueue,
                function(a) {
                    g.debug(a, "\u8fd4\u56deASLB\u2014Data:");
                    f.callback2Player.call(f, e.CALL_PLAYER_TYPE.playAD, a);
                    f.downMaterialTimer = q.now()
                });
                f.sendEvent(e.SEND_EVENT_TYPE.OnAcComplate, {
                    atype: "2",
                    ct: f.dynamicVars.preAdCount
                }); ! 1 === g.isMStation && f.sendEvent(e.SEND_EVENT_TYPE.OnAcComplate, {
                    atype: "3",
                    ct: f.dynamicVars.staAdCount,
                    dur: f.dynamicVars.stadur || "0",
                    oiid: c
                })
            } else f.callback2Player.call(f, e.CALL_PLAYER_TYPE.playAD, []),
            f.sendEvent(e.EVENT_TYPE.OnAcComplate, {
                error: {
                    code: 453
                }
            })
        },
        retry: function(a) {},
        _getUniqueId: function() {
            var a = Math;
            return "ad_" + Array.prototype.join.call(arguments, "_") + String(a.ceil(1E4 * a.random()))
        },
        sendEvent: function(a, c) {
            var b = h;
            try {
                var d = b.config,
                f = d.SEND_EVENT_TYPE,
                e = c.curAD;
                if (e || a == f.OnAcComplate) switch (a) {
                case f.OnAcComplate:
                    b._sendUserLog(0, c);
                    g.debug("AC\u7ed3\u675f");
                    break;
                case f.OnStart:
                    if ("0" == b.dynamicVars.dur_total) break; ! 1 === b.dynamicVars.hasPlayed && (b._sendUserLog(1, c), b._sendArkTracking(1, c), e.sendEvent("start", b._sendArkTracking));
                    0 == b.callback2Player.call(b, d.CALL_PLAYER_TYPE.getRealTime) && (b.playingMonitorCount = b.playingMonitorCount || 0, b.playingMonitor && clearTimeout(b.playingMonitor), b.playingMonitor = setTimeout(function() {++b.playingMonitorCount;
                        5 < b.playingMonitorCount ? b.playingMonitorCount = null: 0 == b.callback2Player.call(b, d.CALL_PLAYER_TYPE.getRealTime) && b.callback2Player.call(b, d.CALL_PLAYER_TYPE.resumeAD)
                    },
                    2E3));
                    b.dynamicVars.hasPlayed = !0;
                    e.adType == d.AD_STYLE.pre_roll ? (e.seeDetail(), e.closeBigPlay(), e.renderRealCd(b.dynamicVars.dur_total, c, b.dynamicVars.dur), b.playAdTimer = b.playAdTimer || [], clearTimeout(b.playAdTimer[c.curIndex]), b.playAdTimer[c.curIndex] = setTimeout(function() {
                        g.debug(c.curIndex + " \u5e7f\u544a\u64ad\u653e\u8d85\u65f6");
                        q.merge(c, {
                            error: {
                                code: 461
                            }
                        });
                        b._sendUserLog(1, c);
                        b.callback2Player.call(b, d.CALL_PLAYER_TYPE.stopAD);
                        e.closeCountDown()
                    },
                    1E3 * e.duration + d.DOWNLOAD_URL_TIMEOUT)) : e.adType == d.AD_STYLE.standard && (e.seeDetail(), e.closeBigPlay(), b.playAdTimer = b.playAdTimer || [], clearTimeout(b.playAdTimer[c.curIndex]), b.playAdTimer[c.curIndex] = setTimeout(function() {
                        g.debug(c.curIndex + " \u5e7f\u544a\u64ad\u653e\u8d85\u65f6");
                        q.merge(c, {
                            error: {
                                code: 461
                            }
                        });
                        b._sendUserLog(1, c);
                        b.callback2Player.call(b, d.CALL_PLAYER_TYPE.stopAD);
                        e.closeSeeDetail()
                    },
                    1E3 * e.duration + d.DOWNLOAD_URL_TIMEOUT));
                    g.debug(c.curIndex + " \u5f00\u59cb\u64ad\u653e\u5e7f\u544a");
                    break;
                case f.OnComplate:
                    e.adType == d.AD_STYLE.pre_roll ? (e.closeSeeDetail(), e.closeBigPlay(c), c.curIndex + 1 == b.dynamicVars.preAdCount ? e.closeCountDown() : e.pauseCountDown()) : e.adType == d.AD_STYLE.standard && e.closeCountDown();
                    clearTimeout(b.playAdTimer[c.curIndex]);
                    b._sendUserLog(3, c);
                    e.sendEvent("complete", b._sendArkTracking);
                    g.debug(c.curIndex + "\u6bb5\u5e7f\u544a\u64ad\u653e\u5b8c\u6210");
                    b.dynamicVars.hasPlayed = !1;
                    break;
                case f.OnPause:
                    b.playingMonitor && clearTimeout(b.playingMonitor);
                    h.playAdTimer && 0 < h.playAdTimer.length && clearTimeout(h.playAdTimer[c.curIndex]);
                    e.pauseCountDown();
                    e.renderBigPlay(c);
                    g.debug(c.curIndex + " \u6682\u505c");
                    break;
                case f.OnError:
                    b._sendUserLog(1, c);
                    if (e.adType == d.AD_STYLE.pre_roll) {
                        clearTimeout(b.playAdTimer[c.curIndex]);
                        for (f = 0; f < c.curIndex; f++);
                        e.closeCountDown()
                    }
                    g.debug(c.error, c.curIndex + " \u64ad\u653e\u5668\u9047\u5230\u9519\u8bef\uff0c\u56de\u8c03");
                    b.dynamicVars.hasPlayed = !1;
                    break;
                case f.OnASLB:
                    b._sendUserLog(5, c);
                    break;
                case f.OnLoginAc:
                    e.loginAc(c.level)
                } else {
                    debugger;
                    b.collectError("1827&err=itemIsNull&type=" + a + "&lc=" + b.putinVars.lc, 3)
                }
            } catch(l) {
                console.log(l),
                b.collectError("974," + (l || {}).stack, 3)
            }
        },
        _sendUserLog: function(a, c) {
            c = c || {};
            var b = h,
            d = b.config,
            f = b.putinVars,
            e = b.dynamicVars,
            t = Math;
            e.dur || q.merge(e, {
                dur: ["-"],
                dur_total: "-",
                adCount: 0
            });
            _adItem = c.curAD || {};
            e = {
                act: "event",
                atype: c.atype || _adItem.adType,
                id: "-",
                ia: 0,
                err: 0,
                lc: f.lc || "-",
                ver: "2.0",
                aps: d.H5_ADPLAYER_VER,
                ch: f.ch,
                cid: f.cid || "-",
                ct: c.ct || 0,
                dur: c.dur || e.dur.join("_") || "0",
                dur_total: c.dur || e.dur_total || "0",
                mmsid: f.mmsid || "-",
                pid: f.pid || "-",
                r: t.ceil(t.random() * q.now()),
                cur_url: encodeURIComponent(location.href),
                ry: e.retry || 0,
                ref: encodeURIComponent(k.referrer) || "-",
                sys: 1,
                uname: f.uname || "-",
                uid: f.uid || "-",
                py: f.up,
                uuid: f.uuid,
                pv: f.ver,
                vid: f.vid || "-",
                vlen: f.gdur || "-",
                p1: f.p1,
                p2: f.p2,
                ontime: "-",
                p3: f.p3 == f.p3 ? "-": f.p3,
                ty: f.islive ? 1 : 0
            };
            switch (a) {
            case 0:
                e.act = "ac";
                e.ry = l.retryCount;
                f.isvip && (e.ia = c.isvip || "1", e.ry = "0");
                c.error && (e.err = c.error.code);
                e.ut = l.responseTime;
                "3" == e.atype && (e.atype = "13");
                e.oiid = c.oiid || b.dynamicVars.adidQueue.join("_") || "-";
                b._sendData(b.config.DC_AD_URL + g.param(e));
                break;
            case 1:
                e.ut = q.now() - b.downMaterialTimer;
                b.lastCostTime = e.ut;
                if (c.error) {
                    switch (c.error.code) {
                    case 1:
                        e.err = 460;
                        break;
                    case 2:
                        e.err = 461;
                        break;
                    case 3:
                        e.err = 463;
                        break;
                    case 4:
                        e.err = 469;
                        break;
                    default:
                        e.err = c.error.code || 0
                    }
                    e.loc = encodeURIComponent(_adItem.url)
                }
                e.dur = _adItem.duration;
                e.ftype = "video";
                e.id = a;
                e.ry = 1;
                e.atype = _adItem.adType;
                e.ord = (parseInt(_adItem.curIdx) || 0) + 1;
                0 < e.ct && e.ord > e.ct && (e.ord = 1, b.collectError("1129&data=" + e.ord + "&idx=" + _adItem.curIndex + "&lc=" + b.putinVars.lc, 3));
                e.atype == d.AD_STYLE.standard ? (e.dur_total = e.dur, e.ord = 1, e.ct = b.dynamicVars.staAdCount) : e.ct = b.dynamicVars.preAdCount;
                "3" == e.atype && (e.atype = "13");
                e.oiid = _adItem.oid || b.dynamicVars.adidQueue[c.curIndex];
                b._sendData(b.config.DC_AD_URL + g.param(e));
                break;
            case 2:
            case 3:
                e.dur = _adItem.duration;
                e.ut = q.now() - b.downMaterialTimer - b.lastCostTime;
                e.ftype = "video";
                e.id = a;
                e.atype = _adItem.adType;
                e.ord = (parseInt(_adItem.curIdx) || 0) + 1;
                0 < e.ct && e.ord > e.ct && (e.ord = 1, b.collectError("1129&data=" + e.ord + "&idx=" + _adItem.curIndex + "&lc=" + b.putinVars.lc, 3));
                e.atype == d.AD_STYLE.standard ? (e.dur_total = e.dur, e.ord = 1, e.ct = b.dynamicVars.staAdCount) : e.ct = b.dynamicVars.preAdCount;
                "3" == e.atype && (e.atype = "13");
                e.oiid = _adItem.oid || b.dynamicVars.adidQueue[_adItem.curIdx];
                b._sendData(b.config.DC_AD_URL + g.param(e));
                3 == a && (b.downMaterialTimer = q.now());
                break;
            case 5:
                g.debug("ASLB\u7ed3\u675f"),
                _adItem.err && (e.loc = encodeURIComponent(_adItem.url), e.err = _adItem.err),
                e.act = "aslb",
                e.ut = _adItem.costTime,
                e.ry = _adItem.ryCount,
                e.atype = _adItem.adType,
                e.ord = (parseInt(_adItem.curIdx) || 0) + 1,
                0 < e.ct && e.ord > e.ct && (e.ord = 1),
                e.atype == d.AD_STYLE.standard && (e.ord = 1, b.collectError("1129&data=" + e.ord + "&idx=" + _adItem.curIndex + "&lc=" + b.putinVars.lc, 3)),
                e.oiid = _adItem.oid || b.dynamicVars.adidQueue[_adItem.curIdx],
                delete e.ct,
                delete e.dur,
                delete e.dur_total,
                delete e.ia,
                "3" == e.atype && (e.atype = "13"),
                b._sendData(b.config.DC_AD_URL + g.param(e))
            }
        },
        _getCtUrl: function(a, c) {
            return this._getAttachParam(a.clickUrl, a.aduid, c || 2, 1, a)
        },
        _getAdStyle: function(a) {
            return this.adStyle ? this.adStyle instanceof Array && this.adStyle.length - 1 >= a ? this.adStyle[a] : this.adStyle: null
        },
        _sendArkTracking: function(a, c, b) {
            var d = [],
            f = c ? c.curAD: {};
            switch (a) {
            case 1:
                d = f.impression;
                for (c = 0; c < d.length; c++) b = "",
                "object" == typeof d[c] ? d[c].cdata && 0 < d[c].cdata.length && (b = d[c].cdata) : b = d[c],
                this._sendData(this._getAttachParam, b, f.aduid, a, 1, f);
                break;
            case 2:
                d = f.tracking;
                for (c = 0; c < d.length; c++) b = "",
                "object" == typeof d[c] ? d[c].cdata && 0 < d[c].cdata.length && (b = d[c].cdata) : b = d[c],
                this._sendData(this._getAttachParam, b, f.aduid, 3, 1, f);
                break;
            case 4:
                if (d = c, f = b, d && 0 < d.length) for (c = 0; c < d.length; c++) this._sendData(this._getAttachParam, d[c], f.aduid, a, 1, f)
            }
        },
        _getAttachParam: function(a, c, b, d, f) {
            var e = h;
            if (!a || "javascript:void(0)" === a) return "javascript:void(0)";
            if ( - 1 < a.indexOf(e.config.ARK_DOMAIN)) {
                var l = (new Date).getTime(),
                k = e.staticVars,
                m = e.putinVars;
                d = {
                    rt: b,
                    oid: f.oid,
                    im: d === s ? 1 : d,
                    t: k.stime + Math.ceil((l - e.startTime) / 1E3),
                    data: [c, k.area_id, k.arkId || 0, m.uuid, f.orderid, m.vid || "", m.pid || "", m.cid || "", f.lc || "1", f.adType || "2", e.putinVars.ch || "letv", g.resoSid(e.putinVars.streamid) || "", f.curIdx + 1 || 0, Math.ceil(l / 1E3), 0, e.putinVars.ver, "", 1].join()
                };
                d.s = e._getSecurityKey(d);
                2 == b ? ( - 1 < a.indexOf("[randnum]") && (a = a.replace("[randnum]", (new Date).getTime())), -1 < a.indexOf("[M_IESID]") && (a = a.replace("[M_IESID]", "LETV_" + c)), -1 < a.indexOf("[M_ADIP]") && (a = a.replace("[M_ADIP]", e.staticVars.ip)), -1 < a.indexOf("[A_ADIP]") && (a = a.replace("[A_ADIP]", e.staticVars.ip)), a = a.split("&u="), a = [a[0], g.param(d), "u=" + a[1]].join("&")) : a += "&" + g.param(d)
            } else - 1 < a.indexOf("[randnum]") && (a = a.replace("[randnum]", (new Date).getTime())),
            -1 < a.indexOf("[M_IESID]") && (a = a.replace("[M_IESID]", "LETV_" + c)),
            -1 < a.indexOf("[M_ADIP]") && (a = a.replace("[M_ADIP]", e.staticVars.ip)),
            -1 < a.indexOf("http://v.admaster.com.cn") && (a = a + ",f" + e.staticVars.ip);
            return a
        },
        _getSecurityKey: function(a) {
            var c = this.config.crc_table,
            b = 0,
            d = 0,
            f = 0,
            e = "",
            g = "",
            h;
            for (h in a) g += a[h];
            e = g.length;
            for (b = 0; b < e; b++) a = g.charCodeAt(b),
            f = c[d & 15 | (a & 15) << 4],
            d = d >> 4 ^ f,
            f = c[d & 15 | a & 240],
            d = d >> 4 ^ f;
            return d.toString(16)
        },
        _sendData: function(a) {
            var c = a;
            "function" == typeof arguments[0] && (c = arguments[0].apply(this, [].slice.call(arguments, 1)));
            if (c && "" != c) {
                var b = q.createElement("img", {
                    src: c
                });
                g.debug("\u53d1\u8d77url : " + c);
                l(b).on("load",
                function() {
                    b = null
                })
            }
        },
        collectError: function(a, c) {
            c = c || 2;
            0 == Math.floor(100 * Math.random()) % c && (a && "object" == typeof a ? (new Image).src = "http://ark.letv.com/apsdbg/?msg=" + encodeURI(a.stack) : (new Image).src = "http://ark.letv.com/apsdbg/?msg=" + encodeURI(a))
        }
    };
    u.prototype = {
        resolveAdParam: function(a) {
            a = g.json(a);
            this.url = a.hdurl && 0 < a.hdurl.length && 960 < g.getDeviceSize.x && 640 < g.getDeviceSize.y ? a.hdurl: a.url;
            if ("1" === a.sg || a.sg === s || !1 === g.isMStation) this.renderCd = !0;
            a.duration && (this.duration = parseInt(a.duration));
            this.pid = a.pid || 0;
            this.vid = a.vid || 0
        },
        initEvent: function() {
            var a = h.config.PROCESS_EVENT_TICKS,
            c, b, d;
            this.progressTicks = [];
            if (this.event && 0 < this.event.length) for (b = 0; b < this.event.length; b++) if (c = this.event[b], c.offset != s) this.progressTicks.push(c.offset);
            else for (d = 0; d < a.length; d++) c.event == a[d].k && (this.event[b].event = "progress", this.event[b].offset = this.duration * a[d].v || 0, this.progressTicks.push(this.event[b].offset));
            u.curAd = this
        },
        sendEvent: function(a, c, b) {
            try {
                var d = this.getTrackArr(a, b);
                c.call(h, 4, d, this)
            } catch(f) {
                g.debug("\u8fdb\u5ea6\u76d1\u6d4b\u51fa\u9519" + f.stack)
            }
        },
        getTrackArr: function(a, c) {
            var b, d = [];
            if (this.event && 0 < this.event.length) for (b = 0; b < this.event.length; b++) this.event[b].event == a && (c != s ? c == this.event[b].offset && (d.push(this.event[b].cdata), this.event[b].event = "hadSent") : (d.push(this.event[b].cdata), this.event[b].event = "hadSent"));
            return d
        },
        renderRealCd: function(a, c, b) {
            var d = this,
            f, e, k = 0,
            p = Math,
            m = h,
            n = g.el("#div_cdown"),
            v = a,
            r;
            0 < d.progressTicks.length ? (clearInterval(d.processTimer), d.processTimer = setInterval(function() {
                r = m.callback2Player(m.config.CALL_PLAYER_TYPE.getRealTime) || 0;
                for (var a = 0; a < d.progressTicks.length; a++) if (1 >= p.abs(d.progressTicks[a] - r)) {
                    g.debug("\u8fdb\u5ea6\u76d1\u6d4b\uff1aoffset:" + d.progressTicks[a] + ",curTime:" + r + "," + a);
                    d.sendEvent("progress", m._sendArkTracking, d.progressTicks[a]);
                    d.progressTicks.splice(a, 1);
                    0 == d.progressTicks.length && clearInterval(d.processTimer);
                    break
                }
            },
            1E3)) : clearInterval(d.processTimer);
            if (g.canBeClicked) if (!0 != d.renderCd) g.removeElem(g.el("#vdo_post_time"));
            else {
                for (f = d.curIdx; 0 <= f; f--) v -= b[f];
                var s = function(a, b, c) {
                    for (f = 0; f < d.curIdx; f++) a -= c[f];
                    k = m.callback2Player(m.config.CALL_PLAYER_TYPE.getRealTime) || 0;
                    return a -= p.ceil(k)
                };
                e = s.apply(this, arguments);
                var u = function(a) {
                    a = a.toString();
                    var b = "",
                    c;
                    for (c = 0; c < a.length; c++) b += '<em id="cd_' + String(c) + '" class="precdImg" style="' + (2 > a.length ? "float:right;": "") + "background-image:url(" + h.config.COUNTDOWN_IMG_URL + ");background-position:0 " + 20 * -parseInt(a.charAt(c)) + 'px;background-repeat: no-repeat;"></em>';
                    return b
                };
                if (g.existEl(n)) n = g.el("#div_cdown");
                else {
                    var n = q.createElement("div", {
                        className: "vdo_post_time",
                        id: "vdo_post_time"
                    }),
                    w = "\u5e7f\u544a ";
                    g.isMStation && (w = "\u5e7f\u544a ");
                    n.innerHTML = ['<div class="vdo_post_rlt"><div class="vdo_time_bg"></div>', '<div class="vdo_time_info"><span id="div_cdown"></span>' + w + "</div>", "</div>"].join("");
                    m.staticVars.countdownElem.appendChild(n);
                    l("#vdo_skip_pre").on("click",
                    function() {
                        d.skipAd()
                    });
                    l("#div_cdown")[0].innerHTML = u(e)
                }
                clearInterval(d.countdownTimer);
                d.countdownTimer = setInterval(function() {
                    var e = s(a, c, b);
                    if (0 > e) d.closeCountDown();
                    else if (e < v) d.pauseCountDown();
                    else {
                        var e = e.toString(),
                        h = a.toString().length - e.length;
                        f = 0;
                        if (0 < h) for (f = 0; f < h; f++) g.el("#cd_" + String(f)).style.backgroundPosition = "0 -200px";
                        for (j = e.length - 1; 0 <= j; j--) {
                            var h = 20 * parseInt(e.charAt(j)),
                            k = g.el("#cd_" + String(j + f));
                            if (g.existEl(k)) k.style.backgroundPosition = "0 " + -h + "px";
                            else {
                                clearInterval(d.countdownTimer);
                                clearInterval(d.processTimer);
                                break
                            }
                        }
                    }
                },
                500)
            }
        },
        renderBigPlay: function(a) {
            if (g.canBeClicked) {
                var c = this,
                b = h.staticVars.countdownElem,
                d = g.el("#btn_a_resume");
                g.existEl(d) && g.removeElem(d);
                d = q.createElement("div", {
                    id: "btn_a_resume",
                    className: "hv_ico_pasued"
                });
                d.style.display = "block";
                b.appendChild(d);
                l(d).on("click",
                function(b) {
                    b.stopPropagation();
                    b.cancelBubble = !0;
                    c.closeBigPlay(a);
                    h.callback2Player(h.config.CALL_PLAYER_TYPE.resumeAD)
                })
            }
        },
        seeDetail: function() {
            if (g.canBeClicked) {
                var a = this,
                c = g.el("#a_see_detail"),
                b = g.el("#a_see_more"),
                d = h.staticVars.countdownElem,
                f = g.el(".hv_ico_pasued"),
                e = h._getCtUrl(a, 2);
                g.existEl(b) ? (g.el("#a_see_detail").setAttribute("href", e), g.el("#a_see_more").setAttribute("href", e)) : (c = q.createElement("a", {
                    target: "_blank",
                    href: e,
                    id: "a_see_detail",
                    className: "aps_mask_cont"
                }), b = q.createElement("div", {
                    className: "vdo_post_detail"
                }), b.innerHTML = ['<div class="vdo_post_rlt"> <div class="vdo_detail_bg"></div>', '<div class="vdo_detail_info"><a id="a_see_more" href="' + e + '" target="_blank">\u4e86\u89e3\u8be6\u60c5<i></i></a></div>', "</div>"].join(""), g.existEl(f) ? (g.isUC || d.insertBefore(c, f), d.insertBefore(b, f)) : (g.isUC || d.appendChild(c), d.appendChild(b)), d = function(b) {
                    b.stopPropagation();
                    b.cancelBubble = !0; ! 1 !== h.dynamicVars.hasPlayed && (e = h._getCtUrl(a, 2), g.el("#a_see_detail").setAttribute("href", e), g.el("#a_see_more").setAttribute("href", e), h.callback2Player(h.config.CALL_PLAYER_TYPE.pauseAD), h._sendUserLog(2, {
                        curAD: a,
                        curIndex: 0
                    }), h._sendArkTracking(2, {
                        curAD: a,
                        curIndex: 0
                    }), a.pid && a.vid && a.openInApp(a.pid, a.vid))
                },
                l(c).on("click", d), l(b).on("click", d))
            }
        },
        openInApp: function(a, c, b) {
            a = "letvclient://msiteAction?actionType=0&pid=" + encodeURIComponent(a) + "&vid=" + encodeURIComponent(c) + "&from=mletv";
            p.Android ? (c = k.createElement("iframe"), c.style.cssText = "width:0px;height:0px;position:fixed;top:0;left:0;", c.src = a, k.body.appendChild(c)) : location.href = a
        },
        closeSeeDetail: function() {
            if (g.canBeClicked) {
                var a = g.el("#a_see_detail"),
                c = g.el(".vdo_post_detail");
                g.existEl(a) && g.removeElem(a);
                g.existEl(c) && g.removeElem(c)
            }
        },
        closeBigPlay: function(a) {
            g.canBeClicked && (a = g.el("#btn_a_resume"), g.existEl(a) && g.removeElem(a))
        },
        closeCountDown: function() {
            if (g.canBeClicked) {
                clearInterval(this.countdownTimer);
                clearInterval(this.processTimer);
                var a = g.el("#vdo_post_time");
                g.existEl(a) && (this.pauseCountDown(), g.removeElem(a));
                this.closeBigPlay();
                this.closeSeeDetail()
            }
        },
        pauseCountDown: function() {
            g.canBeClicked && (clearInterval(this.countdownTimer), clearInterval(this.processTimer))
        },
        skipAd: function() {
            var a = h,
            c = a.config,
            b = c.CALL_PLAYER_TYPE,
            d = a.callback2Player,
            f = a.staticVars.countdownElem,
            e;
            e = g.el(".aps_pop_poster");
            d.call(a, b.pauseAD);
            g.existEl(e) || (e = q.createElement("div", {
                className: "aps_pop_poster",
                id: "aps_login"
            }), e.innerHTML = '<div class="hv_pop_poster"><p class="hv_p1">\u5982\u679c\u60a8\u5df2\u662f\u4f1a\u5458\uff0c\u8bf7\u767b\u5f55</p><p><a href="javascript:;" id="aps_login_button">\u767b\u5f55</a></p><p>\u770b\u5927\u7247\u65e0\u5e7f\u544a\uff0c<span class="hv_org">7\u5929</span>\u4f1a\u5458\u514d\u8d39\u4f53\u9a8c</p><p><a href="http://yuanxian.letv.com/zt2014/7days/index.shtml?ref=H5PADQad" target="_blank" class="blu">\u7acb\u5373\u9886\u53d6</a></p><a href="javascript:;" id="aps_login_close" class="close"><i></i><i class="i_1"></i></a></div>', f.appendChild(e), a._sendData(c.SKIP_AD_CLICK), l("#aps_login_button").on("click",
            function(a) {
                d(b.doLogin)
            }), l("#aps_login_close").on("click",
            function(c) {
                d.call(a, b.resumeAD);
                g.removeElem(g.el("#aps_login"))
            }));
            g.debug("\u70b9\u51fb\u8df3\u8fc7\u5e7f\u544a")
        },
        loginAc: function(a) {
            var c = h,
            b = c.config,
            d = b.CALL_PLAYER_TYPE,
            f = c.callback2Player;
            a ? (f(d.stopAD), this.closeCountDown(), c._sendData(b.SKIP_AD_SUCC)) : f(d.resumeAD);
            g.debug("\u767b\u5f55\u5b8c\u6210\uff0c\u8fd4\u56delevel\uff1a" + a);
            g.removeElem(l(".aps_pop_poster")[0])
        }
    };
    w.prototype = {
        open: function() {
            var a = this;
            a.support ? (ws = new WebSocket(a.target), ws.onopen = function(c) {
                a.onopen.apply(a, arguments)
            },
            ws.onmessage = function(c) {
                a.onmessage.apply(a, arguments)
            },
            ws.onerror = function(c) {
                a.onerror.apply(a, arguments)
            },
            ws.onclose = function() {
                a.onclose.apply(a, arguments)
            },
            a.socket = ws) : alert("your br not support ws");
            a.soldier()
        },
        addLog: function(a) {
            this.mq.push({
                time: +new Date,
                data: a
            })
        },
        send: function(a, c) {
            this.socket.send("[" + c + "]," + a)
        },
        sendHttp: function(a, c) {
            g.sendLogs(a, c)
        },
        close: function() {
            this.support && this.socket.close()
        },
        onopen: function(a) {
            this.ready = !0;
            console.log("onopen:");
            console.log(arguments)
        },
        onmessage: function(a) {
            var c = "",
            b = u.curAd,
            d = h.callback2Player,
            f = h.config.CALL_PLAYER_TYPE;
            if (a.data) switch (c = a.data.split(":")[1].replace("\r", "").replace("\n", ""), c) {
            case "connect":
            case "connected":
                break;
            case "closecd":
                b.closeCountDown();
                break;
            case "closedetail":
                b.closeSeeDetail();
                break;
            case "requestad":
                h.initAD();
                break;
            case "refresh":
                location.reload();
                break;
            case "stopad":
                b.closeCountDown();
                d(f.stopAD);
                break;
            case "resumead":
                d(f.resumeAD);
                break;
            case "pausead":
                d(f.pauseAD);
                break;
            default:
                this.send("error command!")
            }
        },
        onerror: function() {
            alert("ws:error:");
            this.ready = !1
        },
        onclose: function() {
            alert("ws:close:");
            this.ready = !1
        },
        soldier: function() {
            var a = this,
            c = 0,
            b;
            a.mqTimer = setInterval(function() {
                var d, f;
                for (b = a.support && a.ready && a.socket ? a.send: a.sendHttp; d = a.mq.shift();) f = d.data,
                "string" !== typeof f && (f = g.json(f)),
                b.call(a, f, d.time);
                16 < ++c && 0 == a.mq.length && (g.debug("ws: soldier time out!"), clearInterval(a.mqTimer))
            },
            2E3)
        }
    };
    window.H5AD = h
} (document, void 0);