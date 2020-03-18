<template>
    <div class="page-container">
        <md-app>
            <md-app-toolbar class="md-primary">
                <div class="md-toolbar-row">
                    <span class="md-title">My Chat App</span>
                </div>
            </md-app-toolbar>
            <md-app-content>
                <md-field>
                    <label>Message</label>
                    <md-textarea v-model="textarea" disabled v-auto-scroll-bottom></md-textarea>
                </md-field>
                <md-field>
                    <label>Your Message</label>
                    <md-input @keyup.enter ="sendMessage()" v-model="message"></md-input>
                    <md-button class="md-primary md-raised" @click="sendMessage()">Submit</md-button>
                </md-field>
            </md-app-content>
        </md-app>
    </div>
</template>


<script>

//const socket = webSocketFactory.connect("wss://localhost:49677"); // 서버연결


var socket = new WebSocket("ws://localhost:8090/ws")

    export default {
        name: 'Chat',
        created() {
            socket.onmessage = ({data}) => {
                this.textarea += data + "\n"
            }
        },
        data() {
            return {
                textarea: "",
                message: '',
            }

        },
        methods: {
            sendMessage () {
                socket.send(this.message)
                this.message = ''
            }
        }
    }
</script>

<style>
    .md-app {
        height: 800px;
        border: 1px solid rgba(#000, .12);
    }

    .md-textarea {
        height: 300px;
    }
</style>