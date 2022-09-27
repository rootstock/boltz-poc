
export default class BoltzPocService {

    static async CreatePayment(id:number) : Promise<string> {
        const request = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ productId: id })
        };
        return fetch('http://localhost:8080/payment', request)
            .then(response => response.json())
            .then(data => data.lninvoice);
    }

    static async SaveConfig(config:{
        key:string, value:string
    }) : Promise<{Key:string, Value:string}> {
        const request = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(config)
        };
        return fetch('http://localhost:8080/config', request)
            .then(response => response.json())
            .then(data => data);
    }

    static async GetConfig(key?:string) : Promise<{Key:string, Value:string}> {
        return fetch('http://localhost:8080/config/' + key)
        .then((res) => res.json())
    }
}