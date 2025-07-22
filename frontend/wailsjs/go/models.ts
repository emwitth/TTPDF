export namespace main {
	
	export class PDFMapMonster {
	    id: string;
	    name: string;
	    cr: number;
	    monsterType: string;
	    size: string;
	
	    static createFrom(source: any = {}) {
	        return new PDFMapMonster(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.cr = source["cr"];
	        this.monsterType = source["monsterType"];
	        this.size = source["size"];
	    }
	}
	export class PDFMapPage {
	    id: string;
	    pageNumber: number;
	    items: PDFMapMonster[];
	
	    static createFrom(source: any = {}) {
	        return new PDFMapPage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.pageNumber = source["pageNumber"];
	        this.items = this.convertValues(source["items"], PDFMapMonster);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

