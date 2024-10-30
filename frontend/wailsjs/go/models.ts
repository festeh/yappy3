export namespace main {
	
	export class ButtonInfo {
	    text: string;
	    method: string;
	
	    static createFrom(source: any = {}) {
	        return new ButtonInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.text = source["text"];
	        this.method = source["method"];
	    }
	}

}

