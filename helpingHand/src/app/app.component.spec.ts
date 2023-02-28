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
  let htmlElement: HTMLElement;

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
    fixture.detectChanges();
    debugElement = fixture.debugElement.query(By.css('#rightSideBar'))
    htmlElement = debugElement.nativeElement;
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

  it('should display no values at start', () => {
    expect(htmlElement.lastChild?.childNodes.length).toEqual(0);
  })




  
});
