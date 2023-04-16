import { Component, ChangeDetectionStrategy, OnInit, Injectable, ViewEncapsulation } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { lastValueFrom } from 'rxjs'
import { SpinnerService } from './spinner.service'

interface ICurrentHand {
  Val: number
  Suit: string
  Index: number

}

interface ICurrentProb {
  Handname: string
  Prob: number
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class AppComponent implements OnInit {
  public displaySuitVal = false;
  public displayVal = false;
  public darkMode = false;
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
  public currentProb: ICurrentProb[] = []
  public currImgs:string[] = new Array;
  constructor (
    private httpClient: HttpClient,
    public spinnerService: SpinnerService
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
    this.currentProb = await lastValueFrom(this.httpClient.get<ICurrentProb[]>('/api/prob'))
    //Probably a better way to do this with the backend or something
    this.currImgs = [];
    for(let i = 0; i < 7; i++) {
      if(i < this.currentHand.length) {
        this.currImgs.push("../assets/" + (this.currentHand[i].Val + 1) + this.currentHand[i].Suit.toLowerCase() + ".png"); 
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

  async toggleDarkMode() {
    this.darkMode = !this.darkMode;
  }

  async removeAll() {
    // for(let i = 0; i < this.currentHand.length; i++) {
    //    this.Val = 1;
    //    this.Index = i;
    //    this.Suit = "";
    //    this.addCard();
    // }
    console.log("in remove all")
    this.currentHand = await lastValueFrom(this.httpClient.get<ICurrentHand[]>('/api/removeAll'))
    this.currentProb = await lastValueFrom(this.httpClient.get<ICurrentProb[]>('/api/prob'))
    //this.currentProb = await lastValueFrom(this.httpClient.get<ICurrentProb[]>('/api/removeALL'))
    //await this.httpClient.get('/api/removeAll')
  }
 
  async randomizeAll() {
    let suits: string[] = ["Spade", "Club", "Heart", "Diamond"];
    for(let i = 0; i < this.currentHand.length; i++) {
       this.Val = Math.floor(Math.random() * 13) + 1;
       this.Index = i;
       this.Suit = suits[Math.floor(Math.random() * 3)];
       this.addCard();
    }
  }

  async setSuit(_Suit : string) {
    this.displaySuitVal = false;
    if(_Suit === "random") {
      let suits: string[] = ["Spade", "Club", "Heart", "Diamond"];
      this.Suit = suits[Math.floor(Math.random() * 3)];
      this.Val = Math.floor(Math.random() * 13) + 1;
      this.addCard();
    }
    else if (_Suit !== "back") {
      this.Suit = _Suit
      this.displayVal = true;
    }
  }

  async setHand(input : ICurrentHand[]) {
    await this.removeAll();
    console.log(this.currentHand);
    for (let i = 0; i < input.length; i++) {
      this.Index = input[i].Index;
      this.Suit = input[i].Suit
      this.Val = input[i].Val
      this.addCard();
    }
    console.log(this.currentHand);
  }

  async setVal(_Val :number) {
    this.displayVal = false;
    if(_Val !== -1) {
      this.Val = _Val;
      this.addCard();
    }
    else {
      this.displaySuitVal = true;
    }
  }

  async addCard() {
    console.log(this.Suit + " " + this.Val + " " + this.Index)
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
