import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";

export const CHAT_URL = "ws://localhost:8080/setup-ws";
// export const CHAT_URL = "wss://socketsbay.com/wss/v2/1/demo/";

@Injectable()
export class WebsocketService {
  base_pdf_service_api = 'http://localhost:8080/';

  constructor(private http: HttpClient) {}


  createRoom():Observable<any>{
    let url = this.base_pdf_service_api + 'create-room'
    return this.http.get(url)
  }

  joinRoom():Observable<any>{
    let url = this.base_pdf_service_api + 'join-room'
    return this.http.get(url)
  }
}