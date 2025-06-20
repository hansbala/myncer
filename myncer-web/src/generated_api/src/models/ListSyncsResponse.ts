/* tslint:disable */
/* eslint-disable */
/**
 * Myncer API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { Sync } from './Sync';
import {
    SyncFromJSON,
    SyncFromJSONTyped,
    SyncToJSON,
    SyncToJSONTyped,
} from './Sync';

/**
 * 
 * @export
 * @interface ListSyncsResponse
 */
export interface ListSyncsResponse {
    /**
     * 
     * @type {Array<Sync>}
     * @memberof ListSyncsResponse
     */
    syncs: Array<Sync>;
}

/**
 * Check if a given object implements the ListSyncsResponse interface.
 */
export function instanceOfListSyncsResponse(value: object): value is ListSyncsResponse {
    if (!('syncs' in value) || value['syncs'] === undefined) return false;
    return true;
}

export function ListSyncsResponseFromJSON(json: any): ListSyncsResponse {
    return ListSyncsResponseFromJSONTyped(json, false);
}

export function ListSyncsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListSyncsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'syncs': ((json['syncs'] as Array<any>).map(SyncFromJSON)),
    };
}

export function ListSyncsResponseToJSON(json: any): ListSyncsResponse {
    return ListSyncsResponseToJSONTyped(json, false);
}

export function ListSyncsResponseToJSONTyped(value?: ListSyncsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'syncs': ((value['syncs'] as Array<any>).map(SyncToJSON)),
    };
}

