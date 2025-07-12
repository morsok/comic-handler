import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Comic {
  id: number;
  title: string;
  description: string;
}

@Injectable({
  providedIn: 'root'
})
export class ComicService {
  private readonly http = inject(HttpClient);
  private readonly baseUrl = this.getApiBaseUrl();

  private getApiBaseUrl(): string {
    // In development mode (ng serve), call backend directly
    if (this.isDevelopmentMode()) {
      return 'http://127.0.0.1:9999/api';
    }
    // In production, API is served from the same origin
    return '/api';
  }

  private isDevelopmentMode(): boolean {
    return window.location.port === '4200' && window.location.hostname === 'localhost';
  }

  getComics(): Observable<Comic[]> {
    return this.http.get<Comic[]>(`${this.baseUrl}/comics`);
  }

  getComic(id: number): Observable<Comic> {
    return this.http.get<Comic>(`${this.baseUrl}/comics/${id}`);
  }

  createComic(comic: { title: string; description: string }): Observable<Comic> {
    return this.http.post<Comic>(`${this.baseUrl}/comics`, comic);
  }
}
