import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { CHAT_URL, WebsocketService } from '../websocket.service';

@Component({
  selector: 'app-chat-room',
  templateUrl: './chat-room.component.html',
  styleUrls: ['./chat-room.component.scss'],
})
export class ChatRoomComponent {
  receivedMsg: string = '';
  wsSubject: WebSocketSubject<any> = webSocket(CHAT_URL);
  userTextFC = new FormControl('');

  constructor(private wsService: WebsocketService) {}

  ngOnInit() {
    
  }

  setupWS(ev: Event) {
    if (ev) {
      ev.preventDefault();
    }

    this.wsSubject.subscribe({
      next: (msg) => {
        console.log('received from ws server: ', msg);
        this.receivedMsg = msg;
      },
      error: (err) => console.log(err),
      complete: () => console.log('listening complete!'),
    });
  }

  send(ev: Event) {
    if (ev) {
      ev.preventDefault();
    }

    let msg = this.userTextFC.value;
    this.wsSubject.next(msg);
    this.wsSubject.complete();
    console.log('msg sent to ws server!');
    this.userTextFC.reset();
  }

  createRoom(ev: Event) {
    if (ev) {
      ev.preventDefault();
    }

    this.wsService.createRoom().subscribe({});
  }

  joinRoom(ev: Event) {
    if (ev) {
      ev.preventDefault();
    }

    this.wsService.joinRoom().subscribe({});
  }
}
