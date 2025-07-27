export namespace main {
	
	export class AppSettings {
	    defaultPath: string;
	
	    static createFrom(source: any = {}) {
	        return new AppSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.defaultPath = source["defaultPath"];
	    }
	}

}

