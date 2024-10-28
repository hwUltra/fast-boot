export class WebSocketService {
  private socket: WebSocket | null = null;
  private url: string;
  private protocols: string;

  constructor(url: string, protocols: string) {
    this.url = url;
    this.protocols = protocols;
  }

  public connect(): void {
    if (this.socket) {
      console.log("WebSocket已经连接上了");
      return;
    }

    this.socket = new WebSocket(this.url, this.protocols);

    this.socket.onopen = () => {
      console.log("已连接WebSocket");
      // 可以在这里发送一些初始化消息
    };

    this.socket.onmessage = (event) => {
      console.log("WebSocketService 接收到的消息:", event.data);

      // 通过事件派发等方式通知 Vue 组件
      window.dispatchEvent(
        new CustomEvent("onmessageWS", {
          detail: {
            data: event.data,
          },
        })
      );
    };

    this.socket.onerror = (error) => {
      console.error("WebSocket中的错误:", error);
    };

    this.socket.onclose = () => {
      console.log("已断开WebSocket连接");
      //
      window.dispatchEvent(
        new CustomEvent("oncloseWS", {
          detail: {
            data: true,
          },
        })
      );
    };
  }

  public disconnect(): void {
    if (this.socket) {
      this.socket.close();
      this.socket = null;
    }
    window.dispatchEvent(
      new CustomEvent("oncloseWS", {
        detail: {
          data: true,
        },
      })
    );
  }

  // 可以添加其他方法，如发送消息
  public sendMessage(message: string): void {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(message);
    } else {
      console.error("WebSocket未连接");
    }
  }
}
