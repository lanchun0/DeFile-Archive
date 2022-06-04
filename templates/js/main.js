var app = new Vue({
    el:"#app",
    data: {
        loginMsg: "",
        privateKey: "",
        newKey: "",
        registered: false,
        authorizationToken: "N.A.",
        address: "N.A.",
        fileAddr: "N.A.",
        tokenAddr: "N.A.",
        balance: "N.A.",
        asset: {},
        allowance: "N.A.",
        moneyAdd: "",
        addMsg: "N.A.",
        moneyWithdraw: "",
        withdrawMsg: "N.A.",
        moneyApprove: "",
        approveMsg: "N.A.",
        file: "N.A.",
        fileName: "",
        permissionLevel: 'L_0',
        allLv: ['L_0', 'L_1', 'L_2', 'L_3', 'L_4'],
        tradable: false,
        price: "",
        uploadMsg: "N.A.",
        fileID: "N.A.",
        allRow: [],
        extensionIcon: {"txt": "./images/txt.png", "docx": "./images/word.png", "pptx": "./images/ppt.png", "xlsx": "./images/excel.png"},
        queryID: "",
        detailID: "",
        detailOwner: "",
        detailPl: "",
        detailTradable: "",
        detailName: "",
        detailHd: "",
        detailSigner: "",
        detailSize: "",
        detailTs: "",
        queryMsg: "",
        shareAddr: "",
        sharePermission: 'reader',
        allPermission: ['reader', 'writer', 'admin']
    },
    mounted() {
        if (localStorage.balance) {
            this.balance = localStorage.balance
        }
        if (localStorage.fileAddr) {
            this.fileAddr = localStorage.fileAddr
        }
        if (localStorage.tokenAddr) {
            this.tokenAddr = localStorage.tokenAddr
        }
        if (localStorage.authorizationToken) {
            this.authorizationToken = localStorage.authorizationToken
        }
        if (localStorage.allowance) {
            this.allowance = localStorage.allowance
        }
        if (localStorage.address) {
            this.address = localStorage.address
        }
        if (localStorage.detailID) {
            this.detailID = localStorage.detailID
        }
        if (localStorage.detailOwner) {
            this.detailOwner = localStorage.detailOwner
        }
        if (localStorage.detailPl) {
            this.detailPl = localStorage.detailPl
        }
        if (localStorage.detailTradable) {
            this.detailTradable = localStorage.detailTradable
        }
        if (localStorage.detailName) {
            this.detailName = localStorage.detailName
        }
        if (localStorage.detailHd) {
            this.detailHd = localStorage.detailHd
        }
        if (localStorage.detailSigner) {
            this.detailSigner = localStorage.detailSigner
        }
        if (localStorage.detailSize) {
            this.detailSize = localStorage.detailSize
        }
        if (localStorage.detailTs) {
            this.detailTs = localStorage.detailTs
        }
    },
    watch: {
        balance(newBalance) {
            localStorage.balance = newBalance
        },
        fileAddr(newAddr) {
            localStorage.fileAddr = newAddr
        },
        tokenAddr(newAddr) {
            localStorage.tokenAddr = newAddr
        },
        authorizationToken(newToken) {
            localStorage.authorizationToken = newToken
        },
        allowance(newAllow) {
            localStorage.allowance = newAllow
        },
        address(newAddr) {
            localStorage.address = newAddr
        },
        detailID(newID) {
            localStorage.detailID = newID
        },
        detailOwner(newOwner) {
            localStorage.detailOwner = newOwner
        },
        detailPl(newPl) {
            localStorage.detailPl = newPl
        },
        detailTradable(newTradable) {
            localStorage.detailTradable = newTradable
        },
        detailName(newName) {
            localStorage.detailName = newName
        },
        detailHd(newHd) {
            localStorage.detailHd = newHd
        },
        detailSigner(newSign) {
            localStorage.detailSigner = newSign
        },
        detailSize(newSize) {
            localStorage.detailSize = newSize
        },
        detailTs(newTime) {
            localStorage.detailTs = newTime
        },
    },
    methods: {
        registerNew: function (){
            that = this
            axios.post('http://localhost:7051/defile/register', {})
                .then(function (response) {
                    that.newKey = response.data.privatekey
                    that.registered = true
                }, function (err) {
                    alert(err)
                })
        },
        loginSystem: function (){
            that = this
            var bodyFormData = new FormData();
            bodyFormData.append('privatekey', this.privateKey)
            axios({
                method: "post",
                url: "http://localhost:7051/defile/login",
                data: bodyFormData,
                headers: { "Content-Type": "multipart/form-data" },
            })
                .then(function (response) {
                    //handle success
                    that.loginMsg = response.data.msg
                    that.authorizationToken = response.data.token
                    that.balance = response.data.user.balance
                    that.address = response.data.user.address
                    localStorage.setItem('assets',JSON.stringify(response.data.user.assets))
                })
                .catch(function (err) {
                    //handle error
                    alert(err);
                })
            axios.get("http://localhost:7051/view/address")
                .then(function (response) {
                    that.fileAddr = response.data.fileaddr
                    that.tokenAddr = response.data.tokenaddr
                }, function (err) {
                    alert(err)
                })
                setInterval(function (){
                    if(that.loginMsg == "success login") {
                        window.location.href = "userInfo.html"
                    }else if(that.loginMsg == "failed to login") {
                        alert("Invalid private key")
                    }
                }, 20);
        },
        refreshAsset: function () {
            this.asset = JSON.parse(localStorage.getItem('assets'))
        },
        refreshAllowance: function () {
            that = this
            axios({
                method: "get",
                url: "http://localhost:7051/defile/allowance",
                headers: {
                    "Authorization" : "Bearer " + that.authorizationToken
                }
            })
                .then(function (response) {
                    that.allowance = response.data.amount
                }, function (err) {
                    alert(err)
                })
        },
        addMoney: function () {
            that = this
            var bodyFormData = new FormData();
            bodyFormData.append('amount', this.moneyAdd)
            axios({
                method: "post",
                url: "http://localhost:7051/defile/topup",
                data: bodyFormData,
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Authorization" : "Bearer " + that.authorizationToken
                },
            })
                .then(function (response) {
                    if(response.data.msg == "failed to top up") {
                        alert("Failed to top up\n" + response.data.err)
                    }else {
                        alert(response.data.msg + "\nYour top up amount is: " + response.data.amount + "\nYour tx is:\n"
                            + response.data.transaction)
                        that.balance = response.data.user.balance
                        that.moneyAdd = ""
                    }
                }, function (err) {
                    alert(err)
                })
        },
        withdrawMoney: function () {
            that = this
            var bodyFormData = new FormData();
            bodyFormData.append('amount', this.moneyWithdraw)
            axios({
                method: "post",
                url: "http://localhost:7051/defile/withdraw",
                data: bodyFormData,
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Authorization" : "Bearer " + that.authorizationToken
                },
            })
                .then(function (response) {
                    if(response.data.msg == "failed to withdraw") {
                        alert("Failed to withdraw\n" + response.data.err)
                    }else {
                        alert(response.data.msg + "\nYour withdraw amount is: " + response.data.amount
                            + "\nYour tx is:\n" + response.data.transaction)
                        // cheating here
                        that.balance = String(Number(response.data.user.balance) - Number(that.moneyWithdraw))
                        that.moneyWithdraw = ""
                    }
                }, function (err) {
                    alert(err)
                })
        },
        approveMoney: function () {
            that = this
            var bodyFormData = new FormData();
            bodyFormData.append('amount', this.moneyApprove)
            axios({
                method: "post",
                url: "http://localhost:7051/defile/approve",
                data: bodyFormData,
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Authorization" : "Bearer " + that.authorizationToken
                },
            })
                .then(function (response) {
                    if(response.data.msg == "failed to approve") {
                        alert("Failed to approve\n" + response.data.err)
                    }else {
                        alert(response.data.msg + "\nYour approve amount is: " + response.data.price + "\nYour tx is:\n"
                            + response.data.transaction)
                        that.moneyApprove = ""
                    }
                }, function (err) {
                    alert(err)
                })
        },
        logoutSystem: function () {
            window.localStorage.clear();
        },
        onChange: function(event) {
            this.file = event.target.files[0]
            this.fileName = event.target.files[0].name
        },
        submitUp: function () {
            if(this.permissionLevel == "L_0") {
                this.tradable = false
            }
            that = this
            var bodyFormData = new FormData();
            bodyFormData.append('file', this.file)
            bodyFormData.append('permissionlevel', this.permissionLevel)
            bodyFormData.append('tradable', String(this.tradable))
            bodyFormData.append('price', this.price)
            axios({
                method: "post",
                url: "http://localhost:7051/defile/upload",
                data: bodyFormData,
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Authorization" : "Bearer " + that.authorizationToken
                },
            })
                .then(function (response) {
                    if(response.data.msg == "failed to upload a file") {
                        alert(response.data.err)
                    }else {
                        that.fileID = response.data.data.id
                        that.uploadMsg = response.data.msg
                        alert(response.data.msg + "\ntx [created]:\n" + "111"
                            + "\ntx [written]:\n" + "222" + "\nfile permission level is: "
                        + response.data.data.permissionlevel + "\ntradable: " + String(response.data.data.tradable))
                        that.file = "N.A."
                        that.permissionLevel = "L_0"
                        that.tradable = false
                        that.price = ""
                        localStorage.setItem('assets',JSON.stringify(response.data.user.assets))
                    }
                }, function (err) {
                    alert(err)
                })
        },
        allFile: function () {
            this.allRow = []
            that = this
            axios.get("http://localhost:7051/view/all")
                .then(function (response) {
                    if(response.data.data.length > 0) {
                        var x = 0
                        for(var i in response.data.data) {
                            if(x == 0) {
                                that.allRow.push([])
                            }
                            var newOb = response.data.data[i]
                            newOb.extension = newOb.MeteData.filename.split('.').pop().toLowerCase()
                            that.allRow[that.allRow.length - 1].push(newOb)
                            x += 1
                            if(x == 4) {
                                x = 0
                            }
                        }
                    }
                }, function (err){
                    alert(err)
                })
        },
        oneFile: function (fileID) {
            that = this
            axios.get("http://localhost:7051/view/one", {
                params: {
                    id : fileID
                }
                }
            )
                .then(function (response) {
                    if(response.data.msg == "failed to query a file") {
                        alert(response.data.msg)
                    }else {
                        that.detailID = response.data.data.id
                        that.detailOwner = response.data.data.owner
                        that.detailPl = response.data.data.permissionlevel
                        that.detailTradable = String(response.data.data.tradable)
                        that.detailName = response.data.data.MeteData.filename
                        that.detailHd = response.data.data.MeteData.hashdigest
                        that.detailSigner = response.data.data.MeteData.signer
                        that.detailSize = response.data.data.MeteData.size
                        that.detailTs = response.data.data.MeteData.timestamp
                        that.queryMsg = response.data.status
                    }
                }, function (err) {
                    alert(err)
                })
            setInterval(function (){
                if(that.queryMsg == 202) {
                    window.location.href = "detail.html"
                }
            }, 20);
        },
        writeFile: function (fileid){
            that = this
            var bodyFormData = new FormData();
            bodyFormData.append('file', this.file)
            bodyFormData.append('id', fileid)
            axios({
                method: "post",
                url: "http://localhost:7051/defile/write",
                data: bodyFormData,
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Authorization" : "Bearer " + that.authorizationToken
                },
            })
                .then(function (response) {
                    if(response.data.msg == "failed to upload a file") {
                        alert(response.data.err)
                    }else {
                        alert(response.data.msg + "\ntx:\n" + response.data.transaction)
                        that.file = "N.A."
                        that.oneFile(fileid)
                    }
                }, function (err) {
                    alert(err)
                })
        },
        buyFile: function (fileid) {
            that = this
            var bodyFormData = new FormData();
            bodyFormData.append('id', fileid)
            axios({
                method: "post",
                url: "http://localhost:7051/defile/buy",
                data: bodyFormData,
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Authorization" : "Bearer " + that.authorizationToken
                },
            })
                .then(function (response) {
                    if(response.data.msg == "failed to buy a file") {
                        alert(response.data.err)
                    }else {
                        alert(response.data.msg + "\ntx:\n" + response.data.transaction)
                        that.balance = response.data.user.balance
                        localStorage.setItem('assets',JSON.stringify(response.data.user.assets))
                    }
                }, function (err) {
                    alert(err)
                })
        },
        downloadFile: function () {

        },
        shareFile: function (fileid) {
            that = this
            var bodyFormData = new FormData();
            bodyFormData.append('id', fileid)
            bodyFormData.append('to', this.shareAddr)
            bodyFormData.append('permission', this.sharePermission)
            axios({
                method: "post",
                url: "http://localhost:7051/defile/share",
                data: bodyFormData,
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Authorization" : "Bearer " + that.authorizationToken
                },
            })
                .then(function (response) {
                    if(response.data.msg == "failed to share a file") {
                        alert(response.data.err)
                    }else {
                        alert("succeeded in sharing file\n" + "tx is:\n" + response.data.trasaction)
                    }
                }, function (err) {
                    alert(err)
                })
        }
    }
})

var app = new Vue({
    el:"#nav",
    methods: {
        logoutSystem: function () {
            window.localStorage.clear();
        }
    }
})