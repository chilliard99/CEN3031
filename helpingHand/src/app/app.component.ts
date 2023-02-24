import { Component, ChangeDetectionStrategy, OnInit, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { lastValueFrom } from 'rxjs'

interface ICurrentHand {
  Val: number
  Suit: string

}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  public title = 'helpingHand!!!!!';
  public Val = 0
  public Suit = ''
  public currentHand: ICurrentHand[] = [
    {
      Val: 0,
      Suit: 'Test Val'
    }
  ]

  constructor (
    private httpClient: HttpClient
  ) {}

  async ngOnInit() {
    await this.loadCards()
  }

  //This is the problem source here
  //I'm 90% sure it has something to do with the fact
  //That the array has both numbers and strings in it
  //rxjs seems to be the way forward to fix, I'll look into
  //it later. 

  //async loadCards(): Promise<Observable<ICurrentHand[]>> {
  //async loadCards() {
    //this.currentHand = await this.httpClient.get<ICurrentHand[]>('/api/hand')
    //return this.httpClient.get<ICurrentHand[]>('/api/hand')
  //}

  async loadCards() {
    this.currentHand = await lastValueFrom(this.httpClient.get<ICurrentHand[]>('/api/hand'))
    console.log(this.currentHand.length)
    //this.currentHand = await this.httpClient.get<ICurrentHand[]>('/api/hand')
  }

  async addCard() {
    await this.httpClient.post('/api/hand', {
      Suit: this.Suit,
      Val: this.Val
    }).toPromise()
    await this.loadCards()
    this.Suit = ''
    this.Val = 0
  }

}
