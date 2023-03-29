import { Component, ChangeDetectionStrategy, OnInit, Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { lastValueFrom } from 'rxjs'

interface ICurrentHand {
  Val: number
  Suit: string
  Index: number

}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  public displaySuitVal = false;
  public displayVal = false;
  public title = 'helpingHand';
  public Val = 0
  public Suit = ''
  public Index = 0
  // public currentHand: ICurrentHand[] = [
  //   {
  //     Val: 0,
  //     Suit: 'Test Val',
  //     Index: 0
  //   }
  // ]
  public currentHand: ICurrentHand[] = []
  public currImgs:string[] = new Array;
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
    //Probably a better way to do this with the backend or something
    this.currImgs = [];
    for(let i = 0; i < 7; i++) {
      if(i < this.currentHand.length) {
        this.currImgs.push("../assets/" + (this.currentHand[i].Val + 1) + this.currentHand[i].Suit + ".png"); 
      }
      else {
        this.currImgs.push("unfilled");
      }
    }
    console.log(this.currentHand.length)


    //this.currentHand = await this.httpClient.get<ICurrentHand[]>('/api/hand')
  }

  async handleRightClick(_Index: number, _Value: number, _Suit: string, event: any) {
    event.preventDefault();
    this.Index = _Index;
    this.Val = _Value;
    this.Suit = _Suit
    this.addCard();
  }
  
  async displaySuit(_Index: number) {
    this.Index = _Index;
    this.displaySuitVal = !this.displaySuitVal;
  }

  async setSuit(_Suit : string) {
    this.Suit = _Suit
    this.displaySuitVal = false;
    this.displayVal = true;
  }

  async setVal(_Val :number) {
    this.Val = _Val;
    this.displayVal = false;
    this.addCard();
  }

  async addCard() {
    console.log("Suit:" + this.Suit + " Value:" + (Number(this.Val) - 1) + " Index:" + this.Index);
    await this.httpClient.post('/api/hand', {
      Suit: this.Suit,
      Val: Number(this.Val) - 1,
      Index: this.Index
    }).toPromise()
    await this.loadCards()
    this.Suit = ''
    this.Val = 0
    this.Index = 0
  }
  
}
