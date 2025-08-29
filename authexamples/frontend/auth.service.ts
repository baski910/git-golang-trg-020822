currentUserSig = signal<User | undefined | null>(undefined)

createUser(data: any) {
    //console.log(data)
    return this.http.post("https://ktmart.in/api/v2/users/", data, {headers: this.headers})
   }
/*
POST http://{{host}}/api/token HTTP/1.1
content-type: application/json
    
{
    "email": "mukesh@go.com",
    "password": "123465789"
}
*/

authenticate(username: string,password: string): Observable<any>{
    const data = {
      username: username,
      password: password
    }
    console.log(data)
    return this.http.post<any>('https://ktmart.in/api/v1/login/',data, {headers: this.headers})
    .pipe(
      tap((response: any)=>{
        console.log(response.user)
        this.currentUserSig.set(response.user)
        localStorage.setItem('token',response.token)
      })
    )
  }
