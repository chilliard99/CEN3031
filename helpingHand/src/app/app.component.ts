import { Component, ChangeDetectionStrategy, OnInit, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { interval, take, lastValueFrom, Observable } from 'rxjs';

interface ICurrentHand {
  Val: string
  Suit: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  public title = 'helpingHand!!!!!';
  public Val = ''
  public Suit = ''
  public currentHand: ICurrentHand[] = [
    {
      Val: 'Test String',
      Suit: 'Test Val'
    }
  ]

  constructor (
    private httpClient: HttpClient
  ) {}

  async ngOnInit() {
    
  }

  //This is the problem source here
  //I'm 90% sure it has something to do with the fact
  //That the array has both numbers and strings in it
  //rxjs seems to be the way forward to fix, I'll look into
  //it later. 

  async loadCards(): Promise<Observable<ICurrentHand[]>> {
  //async loadCards() {
   // this.currentHand = await this.httpClient.get<ICurrentHand[]>('/api/hand').
    return this.httpClient.get<ICurrentHand[]>('/api/hand')
  }

  async addCard() {
    console.log("hello")
    console.log(this.Suit)
    console.log(this.Val)
    await this.httpClient.post('/api/hand', {
      Suit: this.Suit,
      Val: this.Val
    }).toPromise()
    this.Suit = ''
    this.Val = ''
  }

}
