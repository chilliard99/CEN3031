import { TestBed, ComponentFixture } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';

import { DebugElement, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSlideToggleModule } from '@angular/material/slide-toggle'
import { MatButtonModule } from '@angular/material/button'
import { MatInputModule } from '@angular/material/input'; 
import { MatCardModule } from '@angular/material/card';
import {MatSelectModule} from '@angular/material/select'; 
import {MatTabsModule} from '@angular/material/tabs';
import { FormsModule } from '@angular/forms';
import { lastValueFrom } from 'rxjs'
import { By } from '@angular/platform-browser'

describe('AppComponent', () => {
  let debugElement: DebugElement;
  let fixture: ComponentFixture<AppComponent>;
  let app: AppComponent;
  let rightSideBar: HTMLElement;
  let images: HTMLElement;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
        HttpClientTestingModule,
        BrowserModule,
        AppRoutingModule,
        BrowserAnimationsModule,
        MatSlideToggleModule,
        MatButtonModule,
        HttpClientModule,
        FormsModule,
        MatCardModule,
        MatInputModule,
        MatTabsModule,
        MatSelectModule
      ],
      declarations: [
        AppComponent
      ],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AppComponent);
    app = fixture.componentInstance;
    app.currImgs.push("../assets/3club.png");
    app.currentHand.push({
      Suit: "club",
      Val: 3,
      Index: 0
    });
    fixture.detectChanges();
  });

  //First test to ensure that the app actually gets created
  it('should create the app', () => {
    expect(app).toBeTruthy();
  });

  it('should have cards equal to images', () => {
     const lengthCurrHand = app.currentHand.length;
     const lengthImgHand = app.currImgs.length;
     expect(lengthCurrHand).toEqual(lengthImgHand);
  });

  it('should display new text when a card is added to currentHand', () => {
    const element: DebugElement[] = fixture.debugElement.queryAll(By.css('.rightsideloop'));
    expect(element.length).toBe(1);
  })

  it('should display an image when one is added to currImgs', () => {
     const element: DebugElement[] = fixture.debugElement.queryAll(By.css('.cardImgs'));
     expect(element.length).toBe(1);
  })

  it('should display the proper values of currentHand on the page', () => {
    const element: DebugElement[] = fixture.debugElement.queryAll(By.css('.rightsideloop'));
    element.forEach((obj:DebugElement, index:number) => {
      expect(obj.children[0].nativeElement.innerHTML.trim()).toEqual(app.currentHand[index].Suit + " " + app.currentHand[index].Val.toString());
    })
  })

  it('should display the proper image according to currImage', () => {
    const element =  fixture.debugElement.queryAll(By.css('.cardImgs'));
    element.forEach((obj:DebugElement, index:number) => {
      //Src is shown as http://localhost, meanwhile currImgs shows ..//assets, slice is used to remove the first few dots.
      expect(obj.children[0].children[0].nativeElement.src).toContain((app.currImgs[index]).slice(2));
    })
  })

  it('should display the proper image according to currentHand', () => {
    const element =  fixture.debugElement.queryAll(By.css('.cardImgs'));
    element.forEach((obj:DebugElement, index:number) => {
      expect(obj.children[0].children[0].nativeElement.src).toContain("/assets/" + (app.currentHand[index].Val) + app.currentHand[index].Suit + ".png");
    })
  })


  
});
