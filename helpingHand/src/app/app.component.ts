import { Component, ChangeDetectionStrategy } from '@angular/core';
import { HttpClient } from '@angular/common/http'

interface ICurrentHand {
  Val: number
  Suit: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'helpingHand!!!!!';
  Val = 0
  Suit = ''
  public currentHand: ICurrentHand[] = [
    {
      Val: 0,
      Suit: 'Test Val'
    }
  ]

  constructor (
    private httpClient: HttpClient
  ) {}

  //This is the problem source here
  //I'm 90% sure it has something to do with the fact
  //That the array has both numbers and strings in it
  //rxjs seems to be the way forward to fix, I'll look into
  //it later. 
  /*async loadCards() {
    this.currentHand = await this.httpClient.get<ICurrentHand[]>('/api/hand').toPromise()
  }*/

  async addCard() {
    await this.httpClient.post('/api/hand', {
      Suit: this.Suit,
      Val: this.Val
    }).toPromise()
    this.title = ''
    this.Val = 0
  }

}
